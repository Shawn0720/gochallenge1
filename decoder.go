package drum

import (
	"bufio"
	"bytes"
	"os"
)

// DecodeFile decodes the drum machine file found at the provided path
// and returns a pointer to a parsed pattern which is the entry point to the
// rest of the data.
// TODO: implement
func DecodeFile(path string) (*Pattern, error) {
	ptn := parsePattern(path)
	return &ptn, nil
}

func parsePattern(filePath string) Pattern {
	var ptn Pattern

	// Step 1: read in the file
	f, err := os.Open(filePath)
	check(err)
	defer f.Close()

	// Step 2: set buffer reader
	r := bufio.NewReader(f)

	// Step 3: Parse pattern name from header
	parseHeader(r)
	totalLen := parseLength(r)
	tracksLen := totalLen - patternVersionLength - 4

	ptn.version = parseVersion(r)
	ptn.tempo = parseTempo(r)
	ptn.tracks = parseTracks(r, tracksLen)

	return ptn
}

func printPattern(pattern Pattern) string {
	var buffer bytes.Buffer

	buffer.WriteString(printVersion(pattern.version))
	buffer.WriteString(printTempo(pattern.tempo))

	for _, track := range pattern.tracks {
		buffer.WriteString(printTrack(track))
	}

	return buffer.String()
}

// method for type Pattern
func (ptn Pattern) String() string {
	return printPattern(ptn)
}

// Pattern is the high level representation of the
// drum pattern contained in a .splice file.
// TODO: implement
type Pattern struct {
	version string
	tempo   float32
	tracks  []track
}
