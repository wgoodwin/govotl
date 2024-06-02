package govotl

import (
	"fmt"
	"strings"
)

type VotlDoc []VotlNode

func (vd *VotlDoc) AddChild(line string) error {
	old := *vd
	if strings.HasPrefix(line, "\t") {
		if len(old) == 0 {
			return fmt.Errorf("invalid first line, must not have prefixed tabs")
		}
		err := old[len(old)-1].AddChild(strings.TrimPrefix(line, "\t"))
		if err != nil {
			return err
		}
	}
	switch divineType(line) {
	case VotlTypeHeader:
		// Append a header to the list
	}

	*vd = old
	return nil
}
