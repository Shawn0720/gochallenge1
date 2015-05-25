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
	ptn := parse_pattern(path)
	return &ptn, nil
}

func parse_pattern(filePath string) Pattern {
	var ptn Pattern

	// Step 1: read in the file
	f, err := os.Open(filePath)
	check(err)
	defer f.Close()

	// Step 2: set buffer reader
	r := bufio.NewReader(f)

	// Step 3: Parse pattern name from header
	parse_header(r)
	total_len := parse_length(r)
	tracks_len := total_len - PATTERN_VERSION_LENGTH - 4

	ptn.version = parse_version(r)
	ptn.tempo = parse_tempo(r)
	ptn.tracks = parse_tracks(r, tracks_len)

	return ptn
}

func Print_pattern(pattern Pattern) string {
	var buffer bytes.Buffer

	buffer.WriteString(print_version(pattern.version))
	buffer.WriteString(print_tempo(pattern.tempo))

	for _, track := range pattern.tracks {
		buffer.WriteString(print_track(track))
	}

	return buffer.String()
}

// method for type Pattern
func (this Pattern) String() string {
	return Print_pattern(this)
}

// Pattern is the high level representation of the
// drum pattern contained in a .splice file.
// TODO: implement
type Pattern struct {
	version string
	tempo   float32
	tracks  []Track
}
