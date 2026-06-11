package main

import "testing"

func TestKubus_TableDriven(t *testing.T) {
	tests := []struct {
		name           string
		sisi           float64
		expectedVolume float64
		expectedLuas   float64
	}{
		{"sisi 0", 0, 0, 0},
		{"sisi 1", 1, 1, 6},
		{"sisi 4", 4, 64, 96},
		{"sisi 2.5", 2.5, 15.625, 37.5},
		{"sisi 10", 10, 1000, 600},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := Kubus{tt.sisi}

			if k.Volume() != tt.expectedVolume {
				t.Errorf("Volume(%v) = %v, want %v", tt.sisi, k.Volume(), tt.expectedVolume)
			}

			if k.Luas() != tt.expectedLuas {
				t.Errorf("Luas(%v) = %v, want %v", tt.sisi, k.Luas(), tt.expectedLuas)
			}
		})
	}
}
