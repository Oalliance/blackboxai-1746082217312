package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParticipantRegistration(t *testing.T) {
	router := SetupRouter() // Assuming SetupRouter is accessible here

	payload := map[string]string{
		"name": "Test Participant",
		"type": "Shipper",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/participants", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]interface{}
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Test Participant", response["name"])
	assert.Equal(t, "Shipper", response["type"])
}

func TestCreateFreightQuote(t *testing.T) {
	router := SetupRouter()

	payload := map[string]interface{}{
		"service_category":    "Express",
		"cargo_type":          "General",
		"packaging_mode":      "Box",
		"origin":              "CityA",
		"destination":         "CityB",
		"transportation_mode": "Road",
		"rate":                100.0,
		"valid_until":         "2099-12-31T23:59:59Z",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/quotes", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]interface{}
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "CityA", response["origin"])
	assert.Equal(t, "CityB", response["destination"])
}

func TestTokenMintAndTransfer(t *testing.T) {
	router := SetupRouter()

	// Mint tokens
	mintPayload := map[string]interface{}{
		"participant_id": "participant1",
		"token_id":       "TOKEN1",
		"amount":         1000.0,
	}
	mintBody, _ := json.Marshal(mintPayload)

	mintReq, _ := http.NewRequest("POST", "/tokens/mint", bytes.NewBuffer(mintBody))
	mintReq.Header.Set("Content-Type", "application/json")

	mintRR := httptest.NewRecorder()
	router.ServeHTTP(mintRR, mintReq)

	assert.Equal(t, http.StatusOK, mintRR.Code)

	// Transfer tokens
	transferPayload := map[string]interface{}{
		"from_id":  "participant1",
		"to_id":    "participant2",
		"token_id": "TOKEN1",
		"amount":   200.0,
	}
	transferBody, _ := json.Marshal(transferPayload)

	transferReq, _ := http.NewRequest("POST", "/tokens/transfer", bytes.NewBuffer(transferBody))
	transferReq.Header.Set("Content-Type", "application/json")

	transferRR := httptest.NewRecorder()
	router.ServeHTTP(transferRR, transferReq)

	assert.Equal(t, http.StatusOK, transferRR.Code)
}
