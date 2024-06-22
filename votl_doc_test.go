package govotl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVotlDocAddNode(t *testing.T) {
	var doc VotlDoc

	t.Log("Testing First Add...")
	err := doc.AddNode("Head One")
	if assert.NoError(t, err) {
		assert.Equal(t, 1, len(doc))
		assert.Equal(t, "Head One", doc[0].Text())
	}

	t.Log("Testing Second Add...")
	err = doc.AddNode("Head Two")
	if assert.NoError(t, err) {
		assert.Equal(t, 2, len(doc))
		assert.Equal(t, "Head Two", doc[1].Text())
	}

	t.Log("Testing First Child...")
	err = doc.AddNode("\tChild One")
	if assert.NoError(t, err) {
		assert.Equal(t, 2, len(doc))
		assert.Equal(t, 1, len(doc[1].Children()))
		assert.Equal(t, "Child One", doc[1].Children()[0].Text())
	}

	t.Log("Testing Second Child...")
	err = doc.AddNode("\t\tChild Two")
	if assert.NoError(t, err) {
		assert.Equal(t, 2, len(doc))
		assert.Equal(t, 1, len(doc[1].Children()))
		assert.Equal(t, 1, len(doc[1].Children()[0].Children()))
		assert.Equal(t, "Child Two", doc[1].Children()[0].Children()[0].Text())
	}
}
