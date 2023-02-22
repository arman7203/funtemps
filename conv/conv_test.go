package conv

import (
	"math"
	"testing"
)

func TestFahrenheitToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 32, want: 0},
		{input: 134, want: 56.67},
	}

	for _, tc := range tests {
		got := FahrenheitToCelsius(tc.input)
		if math.Abs(tc.want-got) > 0.01 {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestKelvinToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 273.15, want: 0},
	}

	for _, tc := range tests {
		got := KelvinToCelsius(tc.input)
		if math.Abs(tc.want-got) > 0.01 {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: -273.15, want: 0},
	}

	for _, tc := range tests {
		got := CelsiusToKelvin(tc.input)
		if math.Abs(tc.want-got) > 0.01 {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestCelsiusToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 0, want: 32},
		{input: 56.67, want: 134},
	}

	for _, tc := range tests {
		got := CelsiusToFahrenheit(tc.input)
		if math.Abs(tc.want-got) > 0.01 {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
func TestFahrenheitToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 32, want: 273.15},
		{input: 134, want: 330.372},
	}

	for _, tc := range tests {
		got := FahrenheitToKelvin(tc.input)
		if math.Abs(tc.want-got) > 0.01 {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestKelvinToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 273.15, want: 32},
		{input: 330.372, want: 134},
	}

	for _, tc := range tests {
		got := KelvinToFahrenheit(tc.input)
		if math.Abs(tc.want-got) > 0.01 {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

