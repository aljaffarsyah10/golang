// File untuk keperluan testing(test ataupun benchmark) dipisah dengan file utama, namanya harus berakhiran _test.go, dan package-nya harus sama
// go test 58_UnitTest.go 58_UnitTest_testify_test.go -v
package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	kubus              Kubus   = Kubus{4}
	volumeSeharusnya   float64 = 64
	luasSeharusnya     float64 = 96
	// Asli
	// KelilingSeharusnya float64 = 48
	// Test munculkan pesan error
	kelilingSeharusnya float64 = 60
)

func TestHitungVolume(t *testing.T) {
	assert.Equal(t, kubus.Volume(), volumeSeharusnya, "perhitungan volume salah")
}

func TestHitungLuas(t *testing.T) {
	assert.Equal(t, kubus.Luas(), luasSeharusnya, "perhitungan luas salah")
}

func TestHitungKeliling(t *testing.T) {
	assert.Equal(t, kubus.Keliling(), kelilingSeharusnya, "perhitungan keliling salah")
}