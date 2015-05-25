package drum

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
)

const PATTERN_HEADER_LENGTH = 10
const PATTERN_VERSION_LENGTH = 32

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parse_segment(r *bufio.Reader, len int32) ([]byte, error) {
	buffer := make([]byte, len)
	n, err := r.Read(buffer)

	if int32(n) != len {
		buffer = nil
	}

	buffer = bytes.Trim(buffer, "\x00")

	return buffer, err
}

func parse_length(r *bufio.Reader) int32 {
	var len int32
	err := binary.Read(r, binary.BigEndian, &len)
	check(err)

	return len
}

func parse_header(r *bufio.Reader) string {
	var header string
	buffer, err := parse_segment(r, PATTERN_HEADER_LENGTH)
	check(err)

	header = string(buffer)

	return header
}

func parse_version(r *bufio.Reader) string {

	buffer, err := parse_segment(r, PATTERN_VERSION_LENGTH)
	check(err)

	version := string(buffer)

	return version
}

func parse_tracks(r *bufio.Reader, total_len int32) []Track {
	var tracks []Track
	var temp_track Track
	var temp_track_len int32
	for total_len > 0 {
		temp_track, temp_track_len = parse_track(r)
		tracks = append(tracks, temp_track)

		// TODO:
		//fmt.Print(print_track(temp_track))
		//fmt.Println(total_len)
		//fmt.Println(temp_track_len)

		total_len -= temp_track_len
	}

	fmt.Println(total_len)

	return tracks
}
