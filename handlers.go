package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// SetupRouter sets up HTTP routes for the marketplace and governance
func SetupRouter(marketplace *Marketplace, governance *Governance) *mux.Router {
	router := mux.NewRouter()

	// Input validation middleware
	validateInput := func(next http.HandlerFunc, validateFunc func(r *http.Request) error) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if err := validateFunc(r); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			next(w, r)
		}
	}

	// Validation function for participant registration
	validateParticipant := func(r *http.Request) error {
		var req struct {
			Name string `json:"name"`
			Type string `json:"type"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return err
		}
		if req.Name == "" || req.Type == "" {
			return errors.New("name and type are required")
		}
		return nil
	}

	// Participant routes
	router.HandleFunc("/participants", validateInput(func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Name string `json:"name"`
			Type string `json:"type"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		pType := ParticipantType(req.Type)
		participant := marketplace.RegisterParticipant(req.Name, pType)
		json.NewEncoder(w).Encode(participant)
	}, validateParticipant)).Methods("POST")

	// Freight quote routes
	router.HandleFunc("/quotes", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			ServiceCategory    string  `json:"service_category"`
			CargoType          string  `json:"cargo_type"`
			PackagingMode      string  `json:"packaging_mode"`
			Origin             string  `json:"origin"`
			Destination        string  `json:"destination"`
			TransportationMode string  `json:"transportation_mode"`
			Rate               float64 `json:"rate"`
			ValidUntil         string  `json:"valid_until"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		validUntil, err := time.Parse(time.RFC3339, req.ValidUntil)
		if err != nil {
			http.Error(w, "Invalid date format", http.StatusBadRequest)
			return
		}
		quote := marketplace.CreateFreightQuote(
			ServiceCategory(req.ServiceCategory),
			CargoType(req.CargoType),
			PackagingMode(req.PackagingMode),
			req.Origin,
			req.Destination,
			TransportationMode(req.TransportationMode),
			req.Rate,
			validUntil,
		)
		json.NewEncoder(w).Encode(quote)
	}).Methods("POST")

	// Place bid route
	router.HandleFunc("/bids", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			QuoteID   string  `json:"quote_id"`
			CarrierID string  `json:"carrier_id"`
			BidAmount float64 `json:"bid_amount"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		bid, err := marketplace.PlaceBid(req.QuoteID, req.CarrierID, req.BidAmount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(bid)
	}).Methods("POST")

	// Confirm booking route
	router.HandleFunc("/bookings", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			QuoteID   string `json:"quote_id"`
			BidID     string `json:"bid_id"`
			ShipperID string `json:"shipper_id"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		booking, err := marketplace.ConfirmBooking(req.QuoteID, req.BidID, req.ShipperID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(booking)
	}).Methods("POST")

	// Governance routes
	router.HandleFunc("/proposals", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			ProposerID  string `json:"proposer_id"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		proposal := governance.CreateProposal(req.Title, req.Description, req.ProposerID)
		json.NewEncoder(w).Encode(proposal)
	}).Methods("POST")

	router.HandleFunc("/proposals/{id}/vote", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		proposalID := vars["id"]

		var req struct {
			ParticipantID string `json:"participant_id"`
			Approve      bool   `json:"approve"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		err := governance.VoteProposal(proposalID, req.ParticipantID, req.Approve)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}).Methods("POST")

	// Role assignment route
	router.HandleFunc("/roles/assign", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			UserID string `json:"user_id"`
			Role   string `json:"role"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		role := Role(req.Role)
		if role != AdminRole && role != ShipperRole && role != CarrierRole && role != BrokerRole && role != ForwarderRole {
			http.Error(w, "Invalid role", http.StatusBadRequest)
			return
		}
		marketplace.AccessControl.AssignRole(req.UserID, role)
		w.WriteHeader(http.StatusOK)
	}).Methods("POST")

	// Token routes
	router.HandleFunc("/tokens/mint", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			ParticipantID string  `json:"participant_id"`
			TokenID       string  `json:"token_id"`
			Amount        float64 `json:"amount"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		// Check role
		if !marketplace.AccessControl.CheckRole(req.ParticipantID, AdminRole) {
			http.Error(w, "Unauthorized: Admin role required", http.StatusUnauthorized)
			return
		}
		err := marketplace.SmartContract.MintToken(req.ParticipantID, req.Amount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}).Methods("POST")

	router.HandleFunc("/tokens/transfer", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			FromID  string  `json:"from_id"`
			ToID    string  `json:"to_id"`
			TokenID string  `json:"token_id"`
			Amount  float64 `json:"amount"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		err := marketplace.SmartContract.TransferToken(req.FromID, req.ToID, req.Amount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}).Methods("POST")

	// Dispute routes
	router.HandleFunc("/disputes/raise", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			BookingID string `json:"booking_id"`
			RaiserID  string `json:"raiser_id"`
			Reason    string `json:"reason"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		err := marketplace.SmartContract.RaiseDispute(req.BookingID, req.RaiserID, req.Reason)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}).Methods("POST")

	router.HandleFunc("/disputes/resolve", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			DisputeID  string `json:"dispute_id"`
			ResolverID string `json:"resolver_id"`
			Resolution string `json:"resolution"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		err := marketplace.SmartContract.ResolveDispute(req.DisputeID, req.ResolverID, req.Resolution)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}).Methods("POST")

	// Transport mode specific logic route
	router.HandleFunc("/transport/mode", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			TransportMode string `json:"transport_mode"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		err := marketplace.SmartContract.TransportModeSpecificLogic(req.TransportMode)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}).Methods("POST")

	// Escrow routes
	router.HandleFunc("/escrow/lock", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			ParticipantID string  `json:"participant_id"`
			TokenID       string  `json:"token_id"`
			Amount        float64 `json:"amount"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		err := marketplace.SmartContract.LockTokensInEscrow(req.ParticipantID, req.TokenID, req.Amount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}).Methods("POST")

	router.HandleFunc("/escrow/release", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			ParticipantID string  `json:"participant_id"`
			TokenID       string  `json:"token_id"`
			Amount        float64 `json:"amount"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		err := marketplace.SmartContract.ReleaseEscrowTokens(req.ParticipantID, req.TokenID, req.Amount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}).Methods("POST")

	router.HandleFunc("/escrow/refund", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			ParticipantID string  `json:"participant_id"`
			TokenID       string  `json:"token_id"`
			Amount        float64 `json:"amount"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		err := marketplace.SmartContract.RefundEscrowTokens(req.ParticipantID, req.TokenID, req.Amount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}).Methods("POST")

	// Membership subscription routes
	router.HandleFunc("/membership/subscribe", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			ParticipantID string `json:"participant_id"`
			Type          string `json:"type"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		mType := MembershipType(req.Type)
		if mType != FreightForwarder && mType != CustomsBroker {
			http.Error(w, "Invalid membership type", http.StatusBadRequest)
			return
		}
		membership, err := marketplace.MembershipManager.Subscribe(req.ParticipantID, mType)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(membership)
	}).Methods("POST")

	// Membership status check route
	router.HandleFunc("/membership/status/{participantID}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		participantID := vars["participantID"]
		active := marketplace.MembershipManager.CheckActive(participantID)
		json.NewEncoder(w).Encode(map[string]bool{"active": active})
	}).Methods("GET")

	// Subscription routes
	router.HandleFunc("/subscription/subscribe", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			ParticipantID string `json:"participant_id"`
			Type          string `json:"type"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		mType := MembershipType(req.Type)
		if mType != FreightForwarder && mType != CustomsBroker {
			http.Error(w, "Invalid subscription type", http.StatusBadRequest)
			return
		}
		subscription, err := marketplace.SubscriptionService.Subscribe(req.ParticipantID, mType)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(subscription)
	}).Methods("POST")

	router.HandleFunc("/subscription/status/{participantID}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		participantID := vars["participantID"]
		active := marketplace.SubscriptionService.CheckActive(participantID)
		json.NewEncoder(w).Encode(map[string]bool{"active": active})
	}).Methods("GET")

	return router
}
