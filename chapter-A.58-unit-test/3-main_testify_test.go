package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	kubus              Kubus   = Kubus{4}
	volumeSeharusnya   float64 = 64
	luasSeharusnya     float64 = 96
	kelilingSeharusnya float64 = 48
)

func TestHitungVolume(t *testing.T) {
	assert.Equal(t, volumeSeharusnya, kubus.Volume(), "perhitungan volume salah")
}

func TestHitungLuas(t *testing.T) {
	assert.Equal(t, luasSeharusnya, kubus.Luas(), "perhitungan luas salah")
}

func TestHitungKeliling(t *testing.T) {
	assert.Equal(t, kelilingSeharusnya, kubus.Keliling(), "perhitungan keliling salah")
}
