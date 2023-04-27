package govotl

import (
	"fmt"
	"strings"
)

// ElementType corresponds to the various heading types for an element
// based on their prefixed characters.
type ElementType int64

const (
	Heading  ElementType = iota // The basic heading with no prefix
	Checkbox                    // The checkbox prefix [_] or [X]
)

// VOTLElement corresponds to the basic VOTL Header with corresponding
// children.
type VOTLElement struct {
	Value    string
	Type     ElementType
	Checked  bool // Corresponds to the Checkbox element type
	Children []VOTLElement
}

// AddChild recursively adds children to an element. An error is returned if
// the format doesn't correspond to proper VOTL.
func (ve *VOTLElement) AddChild(e string) error {
	if !strings.HasPrefix(e, "\t") {
		ve.Children = append(ve.Children, NewVOTLElement(e))
		return nil
	}

	if len(ve.Children) == 0 {
		return fmt.Errorf("Malformed header [%s], too man leading tabs.", e)
	}

	child := ve.Children[len(ve.Children)-1]
	if err := child.AddChild(strings.TrimPrefix(e, "\t")); err != nil {
		return err
	}
	ve.Children[len(ve.Children)-1] = child
	return nil
}

// NewVOTLEelement generates an element using the given header. At
// this point it only sets the Value field in the return type to the
// given string, but in future iterations should compile header information
// based on the prefix in the string.
func NewVOTLElement(e string) VOTLElement {
	var result VOTLElement

	result.Type = divineType(e)

	if result.Type == Checkbox {
		result.Checked = strings.HasPrefix(e, "[X]")
	}

	// Trim out the prefix if necessary
	switch result.Type {
	case Checkbox:
		result.Value = stripCheckbox(e, result.Checked)
	case Heading:
		result.Value = e
	}

	return result
}

// Divine the header type based on a set of criteria
func divineType(e string) ElementType {
	if strings.HasPrefix(e, "[_]") || strings.HasPrefix(e, "[X]") {
		return Checkbox
	}

	return Heading
}

// Strip off the checkbox prefix, format dependent
// on whether or not the value is checked ([X])
func stripCheckbox(e string, c bool) string {
	if c {
		return strings.Trim(strings.TrimPrefix(e, "[X]"), " ")
	}
	return strings.Trim(strings.TrimPrefix(e, "[_]"), " ")
}
