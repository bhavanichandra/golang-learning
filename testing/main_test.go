package main

import "testing"

// Table Test
var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"Valid Data", 100.0, 10.0, 10.0, false},
	{"Valid Data", 50.0, 10.0, 5.0, false},
	{"Invalid Data", 100.0, 0.0, 0.0, true},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("Expected an error, but did not get one")
			}
		} else {
			if err != nil {
				t.Error("Did not expect an error, but got one!", err.Error())
			}
		}
		if got != tt.expected {
			t.Errorf("Expected %f but got %f", tt.expected, got)
		}

	}
}
