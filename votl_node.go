package govotl

import (
	"fmt"
	"strings"
)

type VotlNode interface {
	AddChild(string) error
	Children() []VotlNode
	Type() VotlType
	String() string
}

type baseNode struct {
	text     string
	vType    VotlType
	level    int
	children []VotlNode
}

func (bn *baseNode) AddChild(line string) error {
	if strings.HasPrefix(line, "\t") {
		if len(bn.children) == 0 {
			return fmt.Errorf("invalid first line, must not have prefixed tabs")
		}
		err := bn.children[len(bn.children)-1].AddChild(strings.TrimPrefix(line, "\t"))
		if err != nil {
			return err
		}
	}
	switch divineType(line) {
	case VotlTypeHeader:
		// Append a header to the list
	}
	return nil
}

func (bn *baseNode) AddChildNode(node VotlNode) {
	bn.children = append(bn.children, node)
}

func (bn *baseNode) Type() VotlType {
	return bn.vType
}

func (bn *baseNode) appendChildrenString() string {
	line := ""
	for _, n := range bn.children {
		line += "\n" + strings.Repeat("\t", bn.level+1) + n.String()
	}
	return line
}
