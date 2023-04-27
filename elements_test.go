package govotl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivintType(t *testing.T) {
	// Basic Heading
	assert.Equal(t, Heading, divineType("Test Heading"))

	// Checkboxes
	assert.Equal(t, Checkbox, divineType("[_] Unchecked"))
	assert.Equal(t, Checkbox, divineType("[X] Checked"))
}

func TestStripCheckbox(t *testing.T) {
	assert.Equal(t, "Test Heading", stripCheckbox("[_] Test Heading", false))
	assert.Equal(t, "Test Heading", stripCheckbox("[X] Test Heading", true))
}

func TestNewVOTLElement(t *testing.T) {
	headElement := NewVOTLElement("Test Heading")
	assert.Equal(t, Heading, headElement.Type)
	assert.Equal(t, "Test Heading", headElement.Value)

	uncheckedElement := NewVOTLElement("[_] Test Unchecked")
	assert.Equal(t, Checkbox, uncheckedElement.Type)
	assert.False(t, uncheckedElement.Checked)
	assert.Equal(t, "Test Unchecked", uncheckedElement.Value)

	checkedElement := NewVOTLElement("[X] Test Checked")
	assert.Equal(t, Checkbox, checkedElement.Type)
	assert.True(t, checkedElement.Checked)
	assert.Equal(t, "Test Checked", checkedElement.Value)
}

func TestAddChild(t *testing.T) {
	head := NewVOTLElement("Head")

	assert.Error(t, head.AddChild("\tBadSub"))

	if assert.NoError(t, head.AddChild("Sub")) {
		assert.Equal(t, "Sub", head.Children[0].Value)
		if assert.NoError(t, head.AddChild("\tSubSub")) {
			assert.Equal(t, "SubSub", head.Children[0].Children[0].Value)
		}

		if assert.NoError(t, head.AddChild("Sub2")) {
			assert.Equal(t, "Sub2", head.Children[1].Value)
			if assert.NoError(t, head.AddChild("\tSubSub2")) {
				assert.Equal(t, "SubSub2", head.Children[1].Children[0].Value)
			}
		}
	}
}

// TODO Test AddChild
// TODO Test NewVotlElement
// TODO Test StripCheckbox
