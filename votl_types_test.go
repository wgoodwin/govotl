package govotl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivineTypeHeader(t *testing.T) {
	assert.Equal(t, VotlTypeHeader, divineType("Header line"))
}
