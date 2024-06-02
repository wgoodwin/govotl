package govotl

type VotlHeader struct {
	*baseNode
}

func (vh VotlHeader) String() string {
	return vh.text + vh.appendChildrenString()
}
