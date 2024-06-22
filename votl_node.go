package govotl

import (
	"fmt"
	"strings"
)

// Basic VotlNode interface
type VotlNode interface {
	AddChild(string) error
	Children() []VotlNode
	Type() VotlType
	Level(int) // Intention is to implement internally to ensure writing the appropriate number of tabs
	String() string
	Text() string
}

func NewVotlNode(line string) VotlNode {
	switch divineType(line) {
	case VotlTypeHeader:
		return NewVotlHeader(line)
	}
	return nil
}

type baseNode struct {
	text     string
	vType    VotlType
	level    int
	children []VotlNode
}

func (bn *baseNode) Text() string {
	return bn.text
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
		return nil
	}
	bn.AddChildNode(NewVotlNode(line))
	return nil
}

func (bn *baseNode) AddChildNode(node VotlNode) {
	node.Level(bn.level + 1)
	bn.children = append(bn.children, node)
}

func (bn *baseNode) Level(level int) {
	bn.level = level
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

func (bn *baseNode) Children() []VotlNode {
	return bn.children
}
