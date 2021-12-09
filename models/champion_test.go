package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewChampion(t *testing.T) {
	newChampion := NewChampion(
		"38",
		"Kassadin",
		"the Void Walker",
		"Cutting a burning swath through the darkest places of the world, Kassadin knows his days are numbered. A widely traveled Shuriman guide and adventurer, he had chosen to raise a family among the peaceful southern tribes—until the day his village was...",
		)
	expected := &Champion{
		"38",
		"Kassadin",
		"the Void Walker",
		"Cutting a burning swath through the darkest places of the world, Kassadin knows his days are numbered. A widely traveled Shuriman guide and adventurer, he had chosen to raise a family among the peaceful southern tribes—until the day his village was...",
	}

	assert.Equal(t, newChampion, expected)
}