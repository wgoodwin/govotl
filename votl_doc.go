package govotl

import (
	"fmt"
	"strings"
)

type VotlDoc []VotlNode

func (vd *VotlDoc) AddNode(line string) error {
	old := *vd
	if strings.HasPrefix(line, "\t") {
		if len(old) == 0 {
			return fmt.Errorf("invalid first line, must not have prefixed tabs")
		}
		err := old[len(old)-1].AddChild(strings.TrimPrefix(line, "\t"))
		if err != nil {
			return err
		}
	} else { // TODO I expect there's some logic way to handle this without an else statement but I haven't worked that out yet
		old = append(old, NewVotlNode(line))
	}

	*vd = old
	return nil
}
