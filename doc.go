package govotl

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)


// VOTLDoc is a shortcut to keeping the root of all the
// headers in a VOTL document. 
type VOTLDoc []VOTLElement


// NewVOTLDoc reads in a file path parses the document line
// by line into a VOTLDoc, if a read or parsing error occurs,
// an error is returned with
func NewVOTLDoc(sourceFile string) (VOTLDoc, error) {
    var result VOTLDoc

    file, err := os.Open(sourceFile)
    if err != nil {
        return result, fmt.Errorf("failed to open source file [%s]: %s", sourceFile, err.Error())
    }

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    for scanner.Scan() {
        if !strings.HasPrefix(scanner.Text(), "\t") {
            result = append(result, NewVOTLElement(scanner.Text()))
            continue
        }

        if len(result) == 0 {
            return result, fmt.Errorf("Malformed header [%s], too man leading tabs.", scanner.Text())
        }

        child := result[len(result) - 1]
        if err := child.AddChild(strings.TrimPrefix(scanner.Text(), "\t")); err != nil {
            return result, err
        }
        result[len(result) - 1] = child
    }

    return result, nil
}
