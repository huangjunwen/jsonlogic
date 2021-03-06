package jsonlogic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLessCompareStrings(t *testing.T) {
	assert.True(t, less("a", "b"))
}

func TestLessCompareStringAndNil(t *testing.T) {
	assert.False(t, less("a", nil))
}

func TestLessCompareNumberAndNil(t *testing.T) {
	assert.False(t, less(12, nil))
}

func TestEqualsWithBooleans(t *testing.T) {
	assert.True(t, equals(true, true))
}
