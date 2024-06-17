package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLCM0027RemoveElement_1(t *testing.T) {
	assert.Equal(t, removeElement([]int{3, 2, 2, 3}, 3), 2)
}
