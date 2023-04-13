package govotl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewVotlDoc(t *testing.T) {
	//Basic Doc Structure with windows newlines
	var testString string = "Header 1\r\n" +
		"\tSubHeader 1\r\n" +
		"\t\tSubSubHeader 1\r\n" +
		"\t\t\tSubSubSubHeader 1\r\n" +
		"\tSubHeader 2\r\n" +
		"\t\tSubSubHeader 2\r\n" +
		"Header 2\r\n" +
		"\tSubHeader 2\r\n" +
		"Header 3\r\n" +
		"\tSubHeader 3\r\n" +
		"\t\tSubSubHeader 3"

	newDoc, err := NewVOTLDoc(testString)

	if assert.NoError(t, err) {
		// Order matters so we should check them sequentially
		assert.Len(t, newDoc, 3)

		// Check header 1
		h1 := newDoc[0]
		assert.Equal(t, "Header 1", h1.Value)

		// Header 1 children
		assert.Len(t, h1.Children, 2)
		sh1 := h1.Children[0]
		assert.Equal(t, "SubHeader 1", sh1.Value)
		assert.Len(t, sh1.Children, 1)
		assert.Equal(t, "SubSubHeader 1", sh1.Children[0].Value)
		assert.Len(t, sh1.Children[0].Children, 1)
		assert.Equal(t, "SubSubSubHeader 1", sh1.Children[0].Children[0].Value)
		assert.Len(t, sh1.Children[0].Children[0].Children, 0)
		sh2 := h1.Children[1]
		assert.Equal(t, "SubHeader 2", sh2.Value)
		assert.Len(t, sh2.Children, 1)
		assert.Equal(t, "SubSubHeader 2", sh2.Children[0].Value)
		assert.Len(t, sh2.Children[0].Children, 0)

		// Check Header 2
		h2 := newDoc[1]
		assert.Equal(t, "Header 2", h2.Value)
		assert.Len(t, h2.Children, 1)
		assert.Equal(t, "SubHeader 2", h2.Children[0].Value)
		assert.Len(t, h2.Children[0].Children, 0)

		// Check Header 3
		h3 := newDoc[2]
		assert.Equal(t, "Header 3", h3.Value)
		assert.Len(t, h3.Children, 1)
		assert.Equal(t, "SubHeader 3", h3.Children[0].Value)
		assert.Len(t, h3.Children[0].Children, 1)
		assert.Equal(t, "SubSubHeader 3", h3.Children[0].Children[0].Value)
		assert.Len(t, h3.Children[0].Children[0].Children, 0)
	}

}

func TestLoadFile(t *testing.T) {
	newDoc, err := LoadFile("test.otl")

	if assert.NoError(t, err) {
		// Order matters so we should check them sequentially
		assert.Len(t, newDoc, 3)

		// Check header 1
		h1 := newDoc[0]
		assert.Equal(t, "Header 1", h1.Value)

		// Header 1 children
		assert.Len(t, h1.Children, 2)
		sh1 := h1.Children[0]
		assert.Equal(t, "SubHeader 1", sh1.Value)
		assert.Len(t, sh1.Children, 1)
		assert.Equal(t, "SubSubHeader 1", sh1.Children[0].Value)
		assert.Len(t, sh1.Children[0].Children, 1)
		assert.Equal(t, "SubSubSubHeader 1", sh1.Children[0].Children[0].Value)
		assert.Len(t, sh1.Children[0].Children[0].Children, 0)
		sh2 := h1.Children[1]
		assert.Equal(t, "SubHeader 2", sh2.Value)
		assert.Len(t, sh2.Children, 1)
		assert.Equal(t, "SubSubHeader 2", sh2.Children[0].Value)
		assert.Len(t, sh2.Children[0].Children, 0)

		// Check Header 2
		h2 := newDoc[1]
		assert.Equal(t, "Header 2", h2.Value)
		assert.Len(t, h2.Children, 1)
		assert.Equal(t, "SubHeader 2", h2.Children[0].Value)
		assert.Len(t, h2.Children[0].Children, 0)

		// Check Header 3
		h3 := newDoc[2]
		assert.Equal(t, "Header 3", h3.Value)
		assert.Len(t, h3.Children, 1)
		assert.Equal(t, "SubHeader 3", h3.Children[0].Value)
		assert.Len(t, h3.Children[0].Children, 1)
		assert.Equal(t, "SubSubHeader 3", h3.Children[0].Children[0].Value)
		assert.Len(t, h3.Children[0].Children[0].Children, 0)
	}

}
