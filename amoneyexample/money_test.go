package amoneyexample

import (
	"testing"

	"github.com/joaocvr/tdd_golang/domain"
	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := domain.NewDollar(5)

	assert.Equal(t, domain.NewDollar(5), five)
	assert.Equal(t, domain.NewDollar(10), five.Times(2))
	assert.Equal(t, domain.NewDollar(15), five.Times(3))
}
