package calculate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerkalian(t *testing.T) {
	bilA := 10
	bilB := 10
	expected := 100
	actual := Perkalian(bilA, bilB)
	// if actual != expected {
	// 	t.Error("hasil tidak sesuai, actual:", actual, "expected:", expected)
	// }
	assert.Equal(t, expected, actual, "hasil tidak sesuai")
}

func TestPerkalianNegatif(t *testing.T) {
	bilA := -5
	bilB := 10
	expected := -50
	actual := Perkalian(bilA, bilB)
	if actual != expected {
		t.Error("hasil tidak sesuai, actual:", actual, "expected:", expected)
	}
}
