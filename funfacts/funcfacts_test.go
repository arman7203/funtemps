package funfacts

import (
	"testing"
)

func TestGetFunFact(t *testing.T) {
	tests := []struct {
		subject string
		unit    string
		want    string
	}{
		{"sun", "C", "The temperature on the outer layer of the Sun is 5506.85°C.\nThe temperature in the core of the Sun is 15,000,000°C.\n"},
		{"sun", "F", "The temperature on the outer layer of the Sun is 9944.33°F.\nThe temperature in the core of the Sun is 27,000,032°F.\n"},
		{"sun", "K", "The temperature on the outer layer of the Sun is 5780.00K.\nThe temperature in the core of the Sun is 15,027,315K.\n"},
		{"invalid", "C", "Invalid subject. Please use 'sun'.\n"},
		{"sun", "invalid", "Invalid temperature unit. Please use C, F, or K.\n"},
	}

	for _, tc := range tests {
		got := GetFunFact(tc.subject, tc.unit)
		if got != tc.want {
			t.Errorf("GetFunFact(%q, %q) = %q, want %q", tc.subject, tc.unit, got, tc.want)
		}
	}
}
