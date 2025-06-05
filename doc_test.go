package docer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	doc, err := Read("test.json")
	assert.NoError(t, err)
	assert.NotNil(t, doc)
}

func TestDoc_Generate(t *testing.T) {
	doc, err := Read("test.json")
	assert.NoError(t, err)
	assert.NotNil(t, doc)

	err = doc.Generate("test.md")
	assert.NoError(t, err)
}
