package amoneyexample

import (
	"testing"

	"github.com/joaocvr/tdd_golang/domain"
	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := domain.Dollar{}
	five.Amount = 5
	ten := five.Times(2)
	fifteen := five.Times(3)

	assert.Equal(t, int64(5), five.Amount)
	assert.Equal(t, int64(10), ten.Amount)
	assert.Equal(t, int64(15), fifteen.Amount)
}
