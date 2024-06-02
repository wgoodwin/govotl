package govotl

type VotlType string

const (
	VotlTypeHeader VotlType = "header"
)

func divineType(line string) VotlType {
	return VotlTypeHeader
}
