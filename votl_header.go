package govotl

type VotlHeader struct {
	*baseNode
}

func (vh VotlHeader) String() string {
	return vh.text + vh.appendChildrenString()
}

func NewVotlHeader(text string) VotlNode {
	return &VotlHeader{
		baseNode: &baseNode{
			text:  text,
			vType: VotlTypeHeader,
		},
	}
}
