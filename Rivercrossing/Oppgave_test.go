package Rivercrossing

// Package for testing the represention of the state of the riverworld

import "testing"

// Test function implemented in line with the Golang's testing tool
func TestViewState(t *testing.T) {
	wanted := "[kylling korn hs ---\\ \\_rev_/ _________________/---]"
	Putinrev()
	state := ViewState()
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}
