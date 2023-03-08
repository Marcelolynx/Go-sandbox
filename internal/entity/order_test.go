package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIfItGetsAnErrorIfIDIsBlank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "invalid id")
}
