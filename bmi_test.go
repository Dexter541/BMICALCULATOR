package main

import "testing"


func Test_metric_converter(t *testing.T) {
	expected := 26.0
	got := metric_converter(165.54, 71.22)

	if got != expected {
		t.Errorf("Expected: %v Got: %v", expected, got)
	}
}

func Test_standard_converter(t *testing.T) {
	expected := 25.8
	got := standard_converter(5, 7, 165)

	if got != expected {
		t.Errorf("Expected: %v Got: %v", expected, got)
	}
}
func Test_message(t *testing.T) {
	expected := "Overweight"
	got := message(25.8)

	if got != expected {
		t.Errorf("Expected: %v Got: %v", expected, got)
	}
}
