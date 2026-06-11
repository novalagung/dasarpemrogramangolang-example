package main

import "testing"

func TestKategorikanNilai(t *testing.T) {
	tests := []struct {
		name     string
		nilai    int
		expected string
	}{
		{"nilai sempurna", 100, "A"},
		{"nilai A minimum", 90, "A"},
		{"nilai A batas", 89, "B"},
		{"nilai B minimum", 80, "B"},
		{"nilai B batas", 79, "C"},
		{"nilai C minimum", 70, "C"},
		{"nilai C batas", 69, "D"},
		{"nilai D minimum", 60, "D"},
		{"nilai D batas", 59, "E"},
		{"nilai nol", 0, "E"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := KategorikanNilai(tt.nilai)
			if result != tt.expected {
				t.Errorf("KategorikanNilai(%d) = %s, want %s", tt.nilai, result, tt.expected)
			}
		})
	}
}

func TestKategorikanNilai_Parallel(t *testing.T) {
	tests := []struct {
		name     string
		nilai    int
		expected string
	}{
		{"nilai A", 95, "A"},
		{"nilai B", 85, "B"},
		{"nilai C", 75, "C"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := KategorikanNilai(tt.nilai)
			if result != tt.expected {
				t.Errorf("KategorikanNilai(%d) = %s, want %s", tt.nilai, result, tt.expected)
			}
		})
	}
}
