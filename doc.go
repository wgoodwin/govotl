package govotl

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// VOTLDoc is a shortcut to keeping the root of all the
// headers in a VOTL document.
type VOTLDoc []VOTLElement

// LoadFile reads in a file path and returns a parsed VOTLDoc
// if a read or parsing error occurs, we return the error
func LoadFile(sourceFile string) (VOTLDoc, error) {

	sourceString, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open source file [%s]: %s", sourceFile, err.Error())
	}

	return NewVOTLDoc(string(sourceString))
}

// NewVOTLDoc reads in a source string and parses it into a new
// VOTLDoc, if a parsing error occurs, the error is returned
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
