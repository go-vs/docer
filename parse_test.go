package docer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mock struct {
	mockB
	ID string   `json:"id"`
	A  *mockA   `json:"a"`
	AS []*mockA `json:"as"`
}

type mockA struct {
	Name string `json:"name"`
}

type mockB struct {
	Age int `json:"age"`
}

func TestParse(t *testing.T) {
	doc := New().HasBody(mock{}, "json")
	assert.NoError(t, doc.JSON("test.json"))
	assert.NoError(t, doc.Generate("test.md"))
}
