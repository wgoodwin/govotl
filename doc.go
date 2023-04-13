package govotl

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// VOTLDoc is a shortcut to keeping the root of all the
// headers in a VOTL document.
type VOTLDoc []VOTLElement

// LoadFile reads in a file path parses the document line
// by line into a VOTLDoc, if a read or parsing error occurs,
// an error is returned with
func LoadFile(sourceFile string) (VOTLDoc, error) {

	sourceString, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open source file [%s]: %s", sourceFile, err.Error())
	}

	return NewVOTLDoc(string(sourceString))
}

func NewVOTLDoc(source string) (VOTLDoc, error) {
	var result VOTLDoc

	// Split doc string into separate lines
	lines := strings.Split(strings.ReplaceAll(source, "\r\n", "\n"), "\n")

	for _, line := range lines {
		// Filter out whitespace
		if strings.TrimSpace(line) == "" {
			continue
		}

		if !strings.HasPrefix(line, "\t") {
			result = append(result, NewVOTLElement(line))
			continue
		}

		if len(result) == 0 {
			return result, fmt.Errorf("Malformed header [%s], too man leading tabs.", line)
		}

		child := result[len(result)-1]
		if err := child.AddChild(strings.TrimPrefix(line, "\t")); err != nil {
			return result, err
		}
		result[len(result)-1] = child
	}

	return result, nil
}
