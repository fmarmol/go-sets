package sets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetInt(t *testing.T) {

	elems := []int{1, 2, 3, 4}

	s := NewSet(elems...)
	assert.True(t, s.Len() == len(elems))

	for _, elem := range elems {
		assert.True(t, s.Contains(elem))
		s.Remove(elem)
		assert.False(t, s.Contains(elem))
	}
	assert.True(t, s.Len() == 0)
}

func TestSetString(t *testing.T) {
	elems := []string{"1", "2", "3", "4"}

	s := NewSet(elems...)
	assert.True(t, s.Len() == len(elems))

	for _, elem := range elems {
		assert.True(t, s.Contains(elem))
		s.Remove(elem)
		assert.False(t, s.Contains(elem))
	}
	assert.True(t, s.Len() == 0)
}

func TestList(t *testing.T) {
	s := NewSet(1, 2, 3)
	assert.ElementsMatch(t, []int{1, 2, 3}, s.List())
}
