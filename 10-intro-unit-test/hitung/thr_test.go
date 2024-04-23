package hitung

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHitungBonusTHR(t *testing.T) {
	t.Run("test posisi senior", func(t *testing.T) {
		var gaji = 5000000
		posisi := "senior"
		expected := 5000000.00
		actual := HitungBonusTHR(float64(gaji), posisi)
		assert.Equal(t, expected, actual, "output salah")
		assert.NotEqual(t, expected, actual, "output salah")
	})

	t.Run("test posisi c-level", func(t *testing.T) {
		var gaji = 5000000
		posisi := "c-level"
		expected := 12500000.00
		actual := HitungBonusTHR(float64(gaji), posisi)
		assert.Equal(t, expected, actual, "output salah")
	})
	t.Run("posisi manager", func(t *testing.T) {
		var gaji = 5000000
		posisi := "manager"
		expected := 10000000.00
		actual := HitungBonusTHR(float64(gaji), posisi)
		assert.Equal(t, expected, actual, "output salah")
	})

	t.Run("posisi dikosongkan", func(t *testing.T) {
		var gaji = 5000000
		posisi := "junior"
		expected := 0.00
		actual := HitungBonusTHR(float64(gaji), posisi)
		assert.Equal(t, expected, actual, "output salah")
	})
}
