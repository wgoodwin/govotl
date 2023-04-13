package govotl

import (
    "fmt"
    "strings"
)

// VOTLElement corresponds to the basic VOTL Header with corresponding
// children.
type VOTLElement struct {
    Value       string
    Children    []VOTLElement
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

    child := ve.Children[len(ve.Children) - 1]
    if err := child.AddChild(strings.TrimPrefix(e, "\t")); err != nil {
        return err
    }
    ve.Children[len(ve.Children) - 1] = child
    return nil
}


// NewVOTLEelement generates an element using the given header. At
// this point it only sets the Value field in the return type to the
// given string, but in future iterations should compile header information
// based on the prefix in the string.
func NewVOTLElement(e string) VOTLElement {
    return VOTLElement{Value: e}
}

