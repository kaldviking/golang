package main

import "testing"

func TestViewState(t *testing.T) {
	wanted := "Dette er en string"
	state := ViewState()
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}
