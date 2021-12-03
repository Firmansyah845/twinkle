package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindLyric(t *testing.T) {
	param := 4

	resultLyric := algoFindLyric(param)

	switch param {
	case 1:
		assert.Equal(t, "This is Twinkle, twinkle, little star", resultLyric)
	case 2:
		assert.Equal(t, "This is How I wonder what you are and Twinkle, twinkle, little star", resultLyric)
	case 3:
		assert.Equal(t, "This is Up above the world so high and How I wonder what you are and Twinkle, twinkle, little star", resultLyric)
	case 4:
		assert.Equal(t, "This is Like a diamond in the sky and Up above the world so high and How I wonder what you are and Twinkle, twinkle, little star", resultLyric)
	}

}
