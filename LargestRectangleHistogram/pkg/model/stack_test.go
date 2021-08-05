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
	assert.Equal(t, err.Error(), "tried to pop from empty stack")
	assert.Equal(t, 0, val)
}

func TestStackLen(t *testing.T) {
	s := Stack{}

	num_elements_to_pop := 3

	for i := 0; i < num_elements_to_pop; i++ {
		assert.Equal(t, s.Len(), i)
		s.Push(1)
	}

	assert.Equal(t, s.Len(), num_elements_to_pop)

	for i := num_elements_to_pop; i > 0; i-- {
		_, err := s.Pop()
		assert.Nil(t, err)
		assert.Equal(t, s.Len(), i-1)
	}

	assert.Equal(t, s.Len(), 0)
}

func TestStackPeek(t *testing.T) {
	s := Stack{}

	s.Push(0)
	assert.Equal(t, s.Peek(), 0)

	s.Push(1)
	assert.Equal(t, s.Peek(), 1)

	_, err := s.Pop()
	assert.Equal(t, s.Peek(), 0)
	assert.Nil(t, err)
}
