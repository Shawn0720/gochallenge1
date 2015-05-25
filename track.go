package drum

import (
	"bufio"
	"encoding/binary"
)

const TEMPO_SIZE = 4    // in byte
const ID_SIZE = 4       // in byte
const NAME_LEN_SIZE = 1 // in byte
const STEPS_SIZE = 16   // in byte

type Track struct {
	id    int32
	name  string
	steps [16]bool
}

func parse_tempo(r *bufio.Reader) float32 {
	var tempo float32

	err := binary.Read(r, binary.LittleEndian, &tempo)
	check(err)

	return tempo
}

func parse_id(r *bufio.Reader) int32 {
	var id int32
	err := binary.Read(r, binary.LittleEndian, &id)
	check(err)

	return id
}

func parse_name_len(r *bufio.Reader) int32 {
	var len_b byte
	err := binary.Read(r, binary.BigEndian, &len_b)
	check(err)

	return int32(len_b)
}

func parse_name(r *bufio.Reader, len int32) string {
	var name string
	buffer, err := parse_segment(r, len)
	check(err)

	name = string(buffer)

	return name
}

func parse_steps(r *bufio.Reader) [16]bool {
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

func parse_track(r *bufio.Reader) (track Track, track_len int32) {

	track.id = parse_id(r)
	name_len := parse_name_len(r)
	track.name = parse_name(r, name_len)
	track.steps = parse_steps(r)

	track_len = ID_SIZE + NAME_LEN_SIZE + name_len + STEPS_SIZE

	return track, track_len
}
