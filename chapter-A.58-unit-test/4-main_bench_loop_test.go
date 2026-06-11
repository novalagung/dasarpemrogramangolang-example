package main

import "testing"

var kubus = Kubus{4}

func BenchmarkHitungLuasLoop(b *testing.B) {
	for b.Loop() {
		kubus.Luas()
	}
}
