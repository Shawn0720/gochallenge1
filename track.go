package drum

import (
	"bufio"
	"encoding/binary"
)

const tempoSize = 4   // in byte
const idSize = 4      // in byte
const nameLenSize = 1 // in byte
const stepsSize = 16  // in byte

type track struct {
	id    int32
	name  string
	steps [16]bool
}

func parseTempo(r *bufio.Reader) float32 {
	var tempo float32

	err := binary.Read(r, binary.LittleEndian, &tempo)
	check(err)

	return tempo
}

func parseID(r *bufio.Reader) int32 {
	var id int32
	err := binary.Read(r, binary.LittleEndian, &id)
	check(err)

	return id
}

func parseNameLen(r *bufio.Reader) int32 {
	var lenInByte byte
	err := binary.Read(r, binary.BigEndian, &lenInByte)
	check(err)

	return int32(lenInByte)
}

func parseName(r *bufio.Reader, len int32) string {
	var name string
	buffer, err := parseSegment(r, len)
	check(err)

	name = string(buffer)

	return name
}

func parseSteps(r *bufio.Reader) [16]bool {
	var steps [16]bool
	var step byte

	for i := 0; i < 16; i++ {
		err := binary.Read(r, binary.LittleEndian, &step)
		check(err)

		if step == 0 {
			steps[i] = false
		} else {
			steps[i] = true
		}
	}
	return steps
}

func parseTrack(r *bufio.Reader) (trackObj track, trackLen int32) {

	trackObj.id = parseID(r)
	nameLen := parseNameLen(r)
	trackObj.name = parseName(r, nameLen)
	trackObj.steps = parseSteps(r)

	trackLen = idSize + nameLenSize + nameLen + stepsSize

	return trackObj, trackLen
}
