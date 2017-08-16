package main

import (
	"gitlab.com/hokiegeek/biologist"
	"gitlab.com/hokiegeek/life"
	"testing"
)

func TestNewCreateAnalysisResponse(t *testing.T) {
	size := life.Dimensions{Width: 3, Height: 3}
	biologist, err := biologist.NewBiologist(size, life.Blinkers, life.ConwayTester())
	if err != nil {
		t.Fatalf("Unable to create biologist: %s\n", err)
	}

	resp := NewCreateAnalysisResponse(biologist)

	if !resp.Dims.Equals(&size) {
		t.Fatal("Expected size %s but received %s\n", size.String(), resp.Dims.String())
	}
}

/*
func TestNewBiologistUpdateResponse(t *testing.T) {
	size := life.Dimensions{Width: 3, Height: 3}
	biologist, err := life.NewBiologist(size, life.Blinkers, life.ConwayTester())
	if err != nil {
		t.Fatalf("Unable to create biologist: %s\n", err)
	}

	resp := NewBiologistUpdateResponse(biologist, 0, 1)
}
*/

// vim: set foldmethod=marker:
