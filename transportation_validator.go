package main

import (
	"errors"
	"strings"
)

// TransportationValidator validates transport modes and applies specific rules
type TransportationValidator struct {
	allowedModes map[string]bool
}

// NewTransportationValidator creates a new TransportationValidator instance
func NewTransportationValidator() *TransportationValidator {
	return &TransportationValidator{
		allowedModes: map[string]bool{
			"air":       true,
			"sea":       true,
			"road":      true,
			"rail":      true,
			"pipeline":  true,
			"multimodal": true,
		},
	}
}

// ValidateMode checks if the transport mode is allowed
func (tv *TransportationValidator) ValidateMode(mode string) error {
	mode = strings.ToLower(strings.TrimSpace(mode))
	if !tv.allowedModes[mode] {
		return errors.New("invalid transport mode: " + mode)
	}
	return nil
}

// ApplyModeSpecificLogic applies specific rules based on transport mode
func (tv *TransportationValidator) ApplyModeSpecificLogic(mode string) error {
	// Example: add custom logic per mode
	switch strings.ToLower(mode) {
	case "air":
		// Air transport specific rules
	case "sea":
		// Sea transport specific rules
	case "road":
		// Road transport specific rules
	case "rail":
		// Rail transport specific rules
	case "pipeline":
		// Pipeline transport specific rules
	case "multimodal":
		// Multimodal transport specific rules
	default:
		return errors.New("unsupported transport mode: " + mode)
	}
	return nil
}
