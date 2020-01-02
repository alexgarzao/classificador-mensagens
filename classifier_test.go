package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXXX(t *testing.T) {
	msg := "Hello world"
	// cls := classifier.New()
	cls := New()
	assert.Equal(t, []string{}, cls.Types(msg))
}
