package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewID(t *testing.T) {
	id := NewID()
	assert.NotNil(t, id)
}

func TestParseID(t *testing.T) {
	id, err := ParseID("00000000-0000-0000-0000-000000000000")
	assert.Nil(t, err)
	assert.NotNil(t, id)
}

func TestParseIDWithError(t *testing.T) {
	_, err := ParseID("00000000-0000-0000-0000-000000000")
	assert.NotNil(t, err)
}
