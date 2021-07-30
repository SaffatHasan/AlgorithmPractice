package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := Stack{}

	s.Push(1)
	s.Push(2)

	val, err := s.Pop()
	assert.Nil(t, err)
	assert.Equal(t, val, 2)

	val, err = s.Pop()
	assert.Nil(t, err)
	assert.Equal(t, val, 1)

	val, err = s.Pop()
	assert.NotNil(t, err)
	assert.Equal(t, 0, val)
}
