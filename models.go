package main

import "time"

// ParticipantType defines the type of marketplace participant
type ParticipantType string

const (
	Shipper         ParticipantType = "Shipper"
	Consignee       ParticipantType = "Consignee"
	Carrier         ParticipantType = "Carrier"
	FreightForwarder ParticipantType = "FreightForwarder"
	CustomsBroker   ParticipantType = "CustomsBroker"
)

// Participant represents a marketplace participant
type Participant struct {
	ID   string
	Name string
	Type ParticipantType
}

// ServiceCategory defines the logistics service category
type ServiceCategory string

const (
	Import        ServiceCategory = "Import"
	Export        ServiceCategory = "Export"
	Transit       ServiceCategory = "Transit"
	Transshipment ServiceCategory = "Transshipment"
)

type SubCategory string

const (
	// Import subcategories
	ImportAir  SubCategory = "ImportAir"
	ImportSea  SubCategory = "ImportSea"
	ImportLand SubCategory = "ImportLand"

	// Export subcategories
	ExportAir  SubCategory = "ExportAir"
	ExportSea  SubCategory = "ExportSea"
	ExportLand SubCategory = "ExportLand"

	// Transit subcategories
	TransitAir SubCategory = "TransitAir"
	TransitSea SubCategory = "TransitSea"

	// Transshipment subcategories
	TransshipmentAir  SubCategory = "TransshipmentAir"
	TransshipmentSea  SubCategory = "TransshipmentSea"
	TransshipmentLand SubCategory = "TransshipmentLand"
)

type SubCategoryItem string

const (
	// ImportAir items by packaging mode
	ImportAirContainerGeneralCargo SubCategoryItem = "ImportAirContainerGeneralCargo"
	ImportAirLooseGeneralCargo     SubCategoryItem = "ImportAirLooseGeneralCargo"
	ImportAirPalletGeneralCargo    SubCategoryItem = "ImportAirPalletGeneralCargo"
	ImportAirContainerHazardousCargo SubCategoryItem = "ImportAirContainerHazardousCargo"
	ImportAirLooseHazardousCargo     SubCategoryItem = "ImportAirLooseHazardousCargo"
	ImportAirPalletHazardousCargo    SubCategoryItem = "ImportAirPalletHazardousCargo"
	ImportAirContainerBulkCargo     SubCategoryItem = "ImportAirContainerBulkCargo"
	ImportAirLooseBulkCargo         SubCategoryItem = "ImportAirLooseBulkCargo"
	ImportAirPalletBulkCargo        SubCategoryItem = "ImportAirPalletBulkCargo"
	ImportAirContainerLiveCargo     SubCategoryItem = "ImportAirContainerLiveCargo"
	ImportAirLooseLiveCargo         SubCategoryItem = "ImportAirLooseLiveCargo"
	ImportAirPalletLiveCargo        SubCategoryItem = "ImportAirPalletLiveCargo"
	ImportAirContainerLiquidCargo   SubCategoryItem = "ImportAirContainerLiquidCargo"
	ImportAirLooseLiquidCargo       SubCategoryItem = "ImportAirLooseLiquidCargo"
	ImportAirPalletLiquidCargo      SubCategoryItem = "ImportAirPalletLiquidCargo"

	// ImportSea items by packaging mode
	ImportSeaContainerGeneralCargo SubCategoryItem = "ImportSeaContainerGeneralCargo"
	ImportSeaLooseGeneralCargo     SubCategoryItem = "ImportSeaLooseGeneralCargo"
	ImportSeaPalletGeneralCargo    SubCategoryItem = "ImportSeaPalletGeneralCargo"
	ImportSeaContainerHazardousCargo SubCategoryItem = "ImportSeaContainerHazardousCargo"
	ImportSeaLooseHazardousCargo     SubCategoryItem = "ImportSeaLooseHazardousCargo"
	ImportSeaPalletHazardousCargo    SubCategoryItem = "ImportSeaPalletHazardousCargo"
	ImportSeaContainerBulkCargo     SubCategoryItem = "ImportSeaContainerBulkCargo"
	ImportSeaLooseBulkCargo         SubCategoryItem = "ImportSeaLooseBulkCargo"
	ImportSeaPalletBulkCargo        SubCategoryItem = "ImportSeaPalletBulkCargo"
	ImportSeaContainerLCLcargo      SubCategoryItem = "ImportSeaContainerLCLcargo"
	ImportSeaLooseLCLcargo          SubCategoryItem = "ImportSeaLooseLCLcargo"
	ImportSeaPalletLCLcargo         SubCategoryItem = "ImportSeaPalletLCLcargo"
	ImportSeaContainerLiveCargo     SubCategoryItem = "ImportSeaContainerLiveCargo"
	ImportSeaLooseLiveCargo         SubCategoryItem = "ImportSeaLooseLiveCargo"
	ImportSeaPalletLiveCargo        SubCategoryItem = "ImportSeaPalletLiveCargo"
	ImportSeaContainerLiquidCargo   SubCategoryItem = "ImportSeaContainerLiquidCargo"
	ImportSeaLooseLiquidCargo       SubCategoryItem = "ImportSeaLooseLiquidCargo"
	ImportSeaPalletLiquidCargo      SubCategoryItem = "ImportSeaPalletLiquidCargo"

	// ImportLand items by packaging mode
	ImportLandContainerGeneralCargo SubCategoryItem = "ImportLandContainerGeneralCargo"
	ImportLandLooseGeneralCargo     SubCategoryItem = "ImportLandLooseGeneralCargo"
	ImportLandPalletGeneralCargo    SubCategoryItem = "ImportLandPalletGeneralCargo"
	ImportLandContainerHazardousCargo SubCategoryItem = "ImportLandContainerHazardousCargo"
	ImportLandLooseHazardousCargo     SubCategoryItem = "ImportLandLooseHazardousCargo"
	ImportLandPalletHazardousCargo    SubCategoryItem = "ImportLandPalletHazardousCargo"
	ImportLandContainerBulkCargo     SubCategoryItem = "ImportLandContainerBulkCargo"
	ImportLandLooseBulkCargo         SubCategoryItem = "ImportLandLooseBulkCargo"
	ImportLandPalletBulkCargo        SubCategoryItem = "ImportLandPalletBulkCargo"
	ImportLandContainerLCLcargo      SubCategoryItem = "ImportLandContainerLCLcargo"
	ImportLandLooseLCLcargo          SubCategoryItem = "ImportLandLooseLCLcargo"
	ImportLandPalletLCLcargo         SubCategoryItem = "ImportLandPalletLCLcargo"
	ImportLandContainerLiveCargo     SubCategoryItem = "ImportLandContainerLiveCargo"
	ImportLandLooseLiveCargo         SubCategoryItem = "ImportLandLooseLiveCargo"
	ImportLandPalletLiveCargo        SubCategoryItem = "ImportLandPalletLiveCargo"
	ImportLandContainerLiquidCargo   SubCategoryItem = "ImportLandContainerLiquidCargo"
	ImportLandLooseLiquidCargo       SubCategoryItem = "ImportLandLooseLiquidCargo"
	ImportLandPalletLiquidCargo      SubCategoryItem = "ImportLandPalletLiquidCargo"

	// ExportAir items by packaging mode
	ExportAirContainerGeneralCargo  SubCategoryItem = "ExportAirContainerGeneralCargo"
	ExportAirLooseGeneralCargo      SubCategoryItem = "ExportAirLooseGeneralCargo"
	ExportAirPalletGeneralCargo     SubCategoryItem = "ExportAirPalletGeneralCargo"
	ExportAirContainerHazardousCargo  SubCategoryItem = "ExportAirContainerHazardousCargo"
	ExportAirLooseHazardousCargo      SubCategoryItem = "ExportAirLooseHazardousCargo"
	ExportAirPalletHazardousCargo     SubCategoryItem = "ExportAirPalletHazardousCargo"
	ExportAirContainerBulkCargo     SubCategoryItem = "ExportAirContainerBulkCargo"
	ExportAirLooseBulkCargo         SubCategoryItem = "ExportAirLooseBulkCargo"
	ExportAirPalletBulkCargo        SubCategoryItem = "ExportAirPalletBulkCargo"
	ExportAirContainerLiveCargo     SubCategoryItem = "ExportAirContainerLiveCargo"
	ExportAirLooseLiveCargo         SubCategoryItem = "ExportAirLooseLiveCargo"
	ExportAirPalletLiveCargo        SubCategoryItem = "ExportAirPalletLiveCargo"
	ExportAirContainerLiquidCargo   SubCategoryItem = "ExportAirContainerLiquidCargo"
	ExportAirLooseLiquidCargo       SubCategoryItem = "ExportAirLooseLiquidCargo"
	ExportAirPalletLiquidCargo      SubCategoryItem = "ExportAirPalletLiquidCargo"

	// ExportSea items by packaging mode
	ExportSeaContainerGeneralCargo  SubCategoryItem = "ExportSeaContainerGeneralCargo"
	ExportSeaLooseGeneralCargo      SubCategoryItem = "ExportSeaLooseGeneralCargo"
	ExportSeaPalletGeneralCargo     SubCategoryItem = "ExportSeaPalletGeneralCargo"
	ExportSeaContainerHazardousCargo  SubCategoryItem = "ExportSeaContainerHazardousCargo"
	ExportSeaLooseHazardousCargo      SubCategoryItem = "ExportSeaLooseHazardousCargo"
	ExportSeaPalletHazardousCargo     SubCategoryItem = "ExportSeaPalletHazardousCargo"
	ExportSeaContainerBulkCargo     SubCategoryItem = "ExportSeaContainerBulkCargo"
	ExportSeaLooseBulkCargo         SubCategoryItem = "ExportSeaLooseBulkCargo"
	ExportSeaPalletBulkCargo        SubCategoryItem = "ExportSeaPalletBulkCargo"
	ExportSeaContainerLCLcargo      SubCategoryItem = "ExportSeaContainerLCLcargo"
	ExportSeaLooseLCLcargo          SubCategoryItem = "ExportSeaLooseLCLcargo"
	ExportSeaPalletLCLcargo         SubCategoryItem = "ExportSeaPalletLCLcargo"
	ExportSeaContainerLiveCargo     SubCategoryItem = "ExportSeaContainerLiveCargo"
	ExportSeaLooseLiveCargo         SubCategoryItem = "ExportSeaLooseLiveCargo"
	ExportSeaPalletLiveCargo        SubCategoryItem = "ExportSeaPalletLiveCargo"
	ExportSeaContainerLiquidCargo   SubCategoryItem = "ExportSeaContainerLiquidCargo"
	ExportSeaLooseLiquidCargo       SubCategoryItem = "ExportSeaLooseLiquidCargo"
	ExportSeaPalletLiquidCargo      SubCategoryItem = "ExportSeaPalletLiquidCargo"

	// ExportLand items by packaging mode
	ExportLandContainerGeneralCargo SubCategoryItem = "ExportLandContainerGeneralCargo"
	ExportLandLooseGeneralCargo     SubCategoryItem = "ExportLandLooseGeneralCargo"
	ExportLandPalletGeneralCargo    SubCategoryItem = "ExportLandPalletGeneralCargo"
	ExportLandContainerHazardousCargo SubCategoryItem = "ExportLandContainerHazardousCargo"
	ExportLandLooseHazardousCargo     SubCategoryItem = "ExportLandLooseHazardousCargo"
	ExportLandPalletHazardousCargo    SubCategoryItem = "ExportLandPalletHazardousCargo"
	ExportLandContainerBulkCargo     SubCategoryItem = "ExportLandContainerBulkCargo"
	ExportLandLooseBulkCargo         SubCategoryItem = "ExportLandLooseBulkCargo"
	ExportLandPalletBulkCargo        SubCategoryItem = "ExportLandPalletBulkCargo"
	ExportLandContainerLCLcargo      SubCategoryItem = "ExportLandContainerLCLcargo"
	ExportLandLooseLCLcargo          SubCategoryItem = "ExportLandLooseLCLcargo"
	ExportLandPalletLCLcargo         SubCategoryItem = "ExportLandPalletLCLcargo"
	ExportLandContainerLiveCargo     SubCategoryItem = "ExportLandContainerLiveCargo"
	ExportLandLooseLiveCargo         SubCategoryItem = "ExportLandLooseLiveCargo"
	ExportLandPalletLiveCargo        SubCategoryItem = "ExportLandPalletLiveCargo"
	ExportLandContainerLiquidCargo   SubCategoryItem = "ExportLandContainerLiquidCargo"
	ExportLandLooseLiquidCargo       SubCategoryItem = "ExportLandLooseLiquidCargo"
	ExportLandPalletLiquidCargo      SubCategoryItem = "ExportLandPalletLiquidCargo"

	// TransitAir items by packaging mode
	TransitAirContainerGeneralCargo SubCategoryItem = "TransitAirContainerGeneralCargo"
	TransitAirLooseGeneralCargo     SubCategoryItem = "TransitAirLooseGeneralCargo"
	TransitAirPalletGeneralCargo    SubCategoryItem = "TransitAirPalletGeneralCargo"
	TransitAirContainerHazardousCargo SubCategoryItem = "TransitAirContainerHazardousCargo"
	TransitAirLooseHazardousCargo     SubCategoryItem = "TransitAirLooseHazardousCargo"
	TransitAirPalletHazardousCargo    SubCategoryItem = "TransitAirPalletHazardousCargo"
	TransitAirContainerBulkCargo     SubCategoryItem = "TransitAirContainerBulkCargo"
	TransitAirLooseBulkCargo         SubCategoryItem = "TransitAirLooseBulkCargo"
	TransitAirPalletBulkCargo        SubCategoryItem = "TransitAirPalletBulkCargo"
	TransitAirContainerLiveCargo     SubCategoryItem = "TransitAirContainerLiveCargo"
	TransitAirLooseLiveCargo         SubCategoryItem = "TransitAirLooseLiveCargo"
	TransitAirPalletLiveCargo        SubCategoryItem = "TransitAirPalletLiveCargo"
	TransitAirContainerLiquidCargo   SubCategoryItem = "TransitAirContainerLiquidCargo"
	TransitAirLooseLiquidCargo       SubCategoryItem = "TransitAirLooseLiquidCargo"
	TransitAirPalletLiquidCargo      SubCategoryItem = "TransitAirPalletLiquidCargo"

	// TransitSea items by packaging mode
	TransitSeaContainerGeneralCargo SubCategoryItem = "TransitSeaContainerGeneralCargo"
	TransitSeaLooseGeneralCargo     SubCategoryItem = "TransitSeaLooseGeneralCargo"
	TransitSeaPalletGeneralCargo    SubCategoryItem = "TransitSeaPalletGeneralCargo"
	TransitSeaContainerHazardousCargo SubCategoryItem = "TransitSeaContainerHazardousCargo"
	TransitSeaLooseHazardousCargo     SubCategoryItem = "TransitSeaLooseHazardousCargo"
	TransitSeaPalletHazardousCargo    SubCategoryItem = "TransitSeaPalletHazardousCargo"
	TransitSeaContainerBulkCargo     SubCategoryItem = "TransitSeaContainerBulkCargo"
	TransitSeaLooseBulkCargo         SubCategoryItem = "TransitSeaLooseBulkCargo"
	TransitSeaPalletBulkCargo        SubCategoryItem = "TransitSeaPalletBulkCargo"
	TransitSeaContainerLCLcargo      SubCategoryItem = "TransitSeaContainerLCLcargo"
	TransitSeaLooseLCLcargo          SubCategoryItem = "TransitSeaLooseLCLcargo"
	TransitSeaPalletLCLcargo         SubCategoryItem = "TransitSeaPalletLCLcargo"
	TransitSeaContainerLiveCargo     SubCategoryItem = "TransitSeaContainerLiveCargo"
	TransitSeaLooseLiveCargo         SubCategoryItem = "TransitSeaLooseLiveCargo"
	TransitSeaPalletLiveCargo        SubCategoryItem = "TransitSeaPalletLiveCargo"
	TransitSeaContainerLiquidCargo   SubCategoryItem = "TransitSeaContainerLiquidCargo"
	TransitSeaLooseLiquidCargo       SubCategoryItem = "TransitSeaLooseLiquidCargo"
	TransitSeaPalletLiquidCargo      SubCategoryItem = "TransitSeaPalletLiquidCargo"

	// TransshipmentAir items by packaging mode
	TransshipmentAirContainerGeneralCargo SubCategoryItem = "TransshipmentAirContainerGeneralCargo"
	TransshipmentAirLooseGeneralCargo     SubCategoryItem = "TransshipmentAirLooseGeneralCargo"
	TransshipmentAirPalletGeneralCargo    SubCategoryItem = "TransshipmentAirPalletGeneralCargo"
	TransshipmentAirContainerHazardousCargo SubCategoryItem = "TransshipmentAirContainerHazardousCargo"
	TransshipmentAirLooseHazardousCargo     SubCategoryItem = "TransshipmentAirLooseHazardousCargo"
	TransshipmentAirPalletHazardousCargo    SubCategoryItem = "TransshipmentAirPalletHazardousCargo"
	TransshipmentAirContainerBulkCargo     SubCategoryItem = "TransshipmentAirContainerBulkCargo"
	TransshipmentAirLooseBulkCargo         SubCategoryItem = "TransshipmentAirLooseBulkCargo"
	TransshipmentAirPalletBulkCargo        SubCategoryItem = "TransshipmentAirPalletBulkCargo"
	TransshipmentAirContainerLiveCargo     SubCategoryItem = "TransshipmentAirContainerLiveCargo"
	TransshipmentAirLooseLiveCargo         SubCategoryItem = "TransshipmentAirLooseLiveCargo"
	TransshipmentAirPalletLiveCargo        SubCategoryItem = "TransshipmentAirPalletLiveCargo"
	TransshipmentAirContainerLiquidCargo   SubCategoryItem = "TransshipmentAirContainerLiquidCargo"
	TransshipmentAirLooseLiquidCargo       SubCategoryItem = "TransshipmentAirLooseLiquidCargo"
	TransshipmentAirPalletLiquidCargo      SubCategoryItem = "TransshipmentAirPalletLiquidCargo"

	// TransshipmentSea items by packaging mode
	TransshipmentSeaContainerGeneralCargo SubCategoryItem = "TransshipmentSeaContainerGeneralCargo"
	TransshipmentSeaLooseGeneralCargo     SubCategoryItem = "TransshipmentSeaLooseGeneralCargo"
	TransshipmentSeaPalletGeneralCargo    SubCategoryItem = "TransshipmentSeaPalletGeneralCargo"
	TransshipmentSeaContainerHazardousCargo SubCategoryItem = "TransshipmentSeaContainerHazardousCargo"
	TransshipmentSeaLooseHazardousCargo     SubCategoryItem = "TransshipmentSeaLooseHazardousCargo"
	TransshipmentSeaPalletHazardousCargo    SubCategoryItem = "TransshipmentSeaPalletHazardousCargo"
	TransshipmentSeaContainerBulkCargo     SubCategoryItem = "TransshipmentSeaContainerBulkCargo"
	TransshipmentSeaLooseBulkCargo         SubCategoryItem = "TransshipmentSeaLooseBulkCargo"
	TransshipmentSeaPalletBulkCargo        SubCategoryItem = "TransshipmentSeaPalletBulkCargo"
	TransshipmentSeaContainerLCLcargo      SubCategoryItem = "TransshipmentSeaContainerLCLcargo"
	TransshipmentSeaLooseLCLcargo          SubCategoryItem = "TransshipmentSeaLooseLCLcargo"
	TransshipmentSeaPalletLCLcargo         SubCategoryItem = "TransshipmentSeaPalletLCLcargo"
	TransshipmentSeaContainerLiveCargo     SubCategoryItem = "TransshipmentSeaContainerLiveCargo"
	TransshipmentSeaLooseLiveCargo         SubCategoryItem = "TransshipmentSeaLooseLiveCargo"
	TransshipmentSeaPalletLiveCargo        SubCategoryItem = "TransshipmentSeaPalletLiveCargo"
	TransshipmentSeaContainerLiquidCargo   SubCategoryItem = "TransshipmentSeaContainerLiquidCargo"
	TransshipmentSeaLooseLiquidCargo       SubCategoryItem = "TransshipmentSeaLooseLiquidCargo"
	TransshipmentSeaPalletLiquidCargo      SubCategoryItem = "TransshipmentSeaPalletLiquidCargo"

	// TransshipmentLand items by packaging mode
	TransshipmentLandContainerGeneralCargo SubCategoryItem = "TransshipmentLandContainerGeneralCargo"
	TransshipmentLandLooseGeneralCargo     SubCategoryItem = "TransshipmentLandLooseGeneralCargo"
	TransshipmentLandPalletGeneralCargo    SubCategoryItem = "TransshipmentLandPalletGeneralCargo"
	TransshipmentLandContainerHazardousCargo SubCategoryItem = "TransshipmentLandContainerHazardousCargo"
	TransshipmentLandLooseHazardousCargo     SubCategoryItem = "TransshipmentLandLooseHazardousCargo"
	TransshipmentLandPalletHazardousCargo    SubCategoryItem = "TransshipmentLandPalletHazardousCargo"
	TransshipmentLandContainerBulkCargo     SubCategoryItem = "TransshipmentLandContainerBulkCargo"
	TransshipmentLandLooseBulkCargo         SubCategoryItem = "TransshipmentLandLooseBulkCargo"
	TransshipmentLandPalletBulkCargo        SubCategoryItem = "TransshipmentLandPalletBulkCargo"
	TransshipmentLandContainerLCLcargo      SubCategoryItem = "TransshipmentLandContainerLCLcargo"
	TransshipmentLandLooseLCLcargo          SubCategoryItem = "TransshipmentLandLooseLCLcargo"
	TransshipmentLandPalletLCLcargo         SubCategoryItem = "TransshipmentLandPalletLCLcargo"
	TransshipmentLandContainerLiveCargo     SubCategoryItem = "TransshipmentLandContainerLiveCargo"
	TransshipmentLandLooseLiveCargo         SubCategoryItem = "TransshipmentLandLooseLiveCargo"
	TransshipmentLandPalletLiveCargo        SubCategoryItem = "TransshipmentLandPalletLiveCargo"
	TransshipmentLandContainerLiquidCargo   SubCategoryItem = "TransshipmentLandContainerLiquidCargo"
	TransshipmentLandLooseLiquidCargo       SubCategoryItem = "TransshipmentLandLooseLiquidCargo"
	TransshipmentLandPalletLiquidCargo      SubCategoryItem = "TransshipmentLandPalletLiquidCargo"
)

// ValidateServiceCategory validates the service category and its relation to freight quote and bidding
func ValidateServiceCategory(category ServiceCategory, transportationMode TransportationMode, originCode, destinationCode string) error {
	switch category {
	case Import, Export:
		// For Import and Export, origin and destination codes must be valid for the transportation mode
		if transportationMode == Air {
			if !validIATAAirportCodes[originCode] {
				return errors.New("origin code is not a valid IATA airport code for Import/Export")
			}
			if !validIATAAirportCodes[destinationCode] {
				return errors.New("destination code is not a valid IATA airport code for Import/Export")
			}
		} else if transportationMode == Sea {
			if !validIMOSeraportCodes[originCode] {
				return errors.New("origin code is not a valid IMO seaport code for Import/Export")
			}
			if !validIMOSeraportCodes[destinationCode] {
				return errors.New("destination code is not a valid IMO seaport code for Import/Export")
			}
		} else {
			return errors.New("unsupported transportation mode for Import/Export")
		}
	case Transit:
		// Transit can only be Air or Sea
		if transportationMode != Air && transportationMode != Sea {
			return errors.New("Transit service category can only be Air or Sea transportation mode")
		}
		// Validate origin and destination codes
		if transportationMode == Air {
			if !validIATAAirportCodes[originCode] {
				return errors.New("origin code is not a valid IATA airport code for Transit")
			}
			if !validIATAAirportCodes[destinationCode] {
				return errors.New("destination code is not a valid IATA airport code for Transit")
			}
		} else if transportationMode == Sea {
			if !validIMOSeraportCodes[originCode] {
				return errors.New("origin code is not a valid IMO seaport code for Transit")
			}
			if !validIMOSeraportCodes[destinationCode] {
				return errors.New("destination code is not a valid IMO seaport code for Transit")
			}
		}
	case Transshipment:
		// Transshipment validation can be similar to Transit or customized as needed
		// For now, allow all transportation modes and validate codes accordingly
		if transportationMode == Air {
			if !validIATAAirportCodes[originCode] {
				return errors.New("origin code is not a valid IATA airport code for Transshipment")
			}
			if !validIATAAirportCodes[destinationCode] {
				return errors.New("destination code is not a valid IATA airport code for Transshipment")
			}
		} else if transportationMode == Sea {
			if !validIMOSeraportCodes[originCode] {
				return errors.New("origin code is not a valid IMO seaport code for Transshipment")
			}
			if !validIMOSeraportCodes[destinationCode] {
				return errors.New("destination code is not a valid IMO seaport code for Transshipment")
			}
		} else if transportationMode == Land {
			// Land transportation validation can be added here if needed
		} else {
			return errors.New("unsupported transportation mode for Transshipment")
		}
	default:
		return errors.New("unknown service category")
	}
	return nil
}

// ValidateServiceCategory validates the service category and its relation to freight quote and bidding
func ValidateServiceCategory(category ServiceCategory, transportationMode TransportationMode, originCode, destinationCode string) error {
	switch category {
	case Import, Export:
		// For Import and Export, origin and destination codes must be valid for the transportation mode
		if transportationMode == Air {
			if !validIATAAirportCodes[originCode] {
				return errors.New("origin code is not a valid IATA airport code for Import/Export")
			}
			if !validIATAAirportCodes[destinationCode] {
				return errors.New("destination code is not a valid IATA airport code for Import/Export")
			}
		} else if transportationMode == Sea {
			if !validIMOSeraportCodes[originCode] {
				return errors.New("origin code is not a valid IMO seaport code for Import/Export")
			}
			if !validIMOSeraportCodes[destinationCode] {
				return errors.New("destination code is not a valid IMO seaport code for Import/Export")
			}
		} else {
			return errors.New("unsupported transportation mode for Import/Export")
		}
	case Transit:
		// Transit can only be Air or Sea
		if transportationMode != Air && transportationMode != Sea {
			return errors.New("Transit service category can only be Air or Sea transportation mode")
		}
		// Validate origin and destination codes
		if transportationMode == Air {
			if !validIATAAirportCodes[originCode] {
				return errors.New("origin code is not a valid IATA airport code for Transit")
			}
			if !validIATAAirportCodes[destinationCode] {
				return errors.New("destination code is not a valid IATA airport code for Transit")
			}
		} else if transportationMode == Sea {
			if !validIMOSeraportCodes[originCode] {
				return errors.New("origin code is not a valid IMO seaport code for Transit")
			}
			if !validIMOSeraportCodes[destinationCode] {
				return errors.New("destination code is not a valid IMO seaport code for Transit")
			}
		}
	case Transshipment:
		// Transshipment validation can be similar to Transit or customized as needed
		// For now, allow all transportation modes and validate codes accordingly
		if transportationMode == Air {
			if !validIATAAirportCodes[originCode] {
				return errors.New("origin code is not a valid IATA airport code for Transshipment")
			}
			if !validIATAAirportCodes[destinationCode] {
				return errors.New("destination code is not a valid IATA airport code for Transshipment")
			}
		} else if transportationMode == Sea {
			if !validIMOSeraportCodes[originCode] {
				return errors.New("origin code is not a valid IMO seaport code for Transshipment")
			}
			if !validIMOSeraportCodes[destinationCode] {
				return errors.New("destination code is not a valid IMO seaport code for Transshipment")
			}
		} else if transportationMode == Land {
			// Land transportation validation can be added here if needed
		} else {
			return errors.New("unsupported transportation mode for Transshipment")
		}
	default:
		return errors.New("unknown service category")
	}
	return nil
}

// CargoType defines the type of cargo
type CargoType string

const (
	GeneralCargo CargoType = "GeneralCargo"
	Perishable  CargoType = "Perishable"
	Hazardous   CargoType = "Hazardous"
	Fragile     CargoType = "Fragile"
)

// PackagingMode defines the packaging mode
type PackagingMode string

const (
	Container PackagingMode = "Container"
	Loose     PackagingMode = "Loose"
	Pallet    PackagingMode = "Pallet"
)

// TransportationMode defines mode of transportation
type TransportationMode string

const (
	Sea TransportationMode = "Sea"
	Air TransportationMode = "Air"
	Land TransportationMode = "Land"
)

// FreightQuote represents a freight quotation
type FreightQuote struct {
	ID                 string
	ServiceCategory    ServiceCategory
	CargoType          CargoType
	PackagingMode      PackagingMode
	OriginCode         string // IATA airport code or IMO seaport code
	DestinationCode    string // IATA airport code or IMO seaport code
	TransportationMode TransportationMode
	Rate               float64
	ValidUntil         time.Time
}

// FreightBid represents a bid on a freight quote
type FreightBid struct {
	ID          string
	QuoteID     string
	CarrierID   string
	BidAmount   float64
	BidTime     time.Time
	IsAccepted  bool
}

// Booking represents a confirmed cargo booking
type Booking struct {
	ID          string
	QuoteID     string
	BidID       string
	ShipperID   string
	CarrierID   string
	BookingTime time.Time
	Status      string
}
