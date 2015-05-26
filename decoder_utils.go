package drum

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
)

const patternHeaderLength = 10
const patternVersionLength = 32

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseSegment(r *bufio.Reader, len int32) ([]byte, error) {
	buffer := make([]byte, len)
	n, err := r.Read(buffer)

	if int32(n) != len {
		buffer = nil
	}

	buffer = bytes.Trim(buffer, "\x00")

	return buffer, err
}

func parseLength(r *bufio.Reader) int32 {
	var len int32
	err := binary.Read(r, binary.BigEndian, &len)
	check(err)

	return len
}

func parseHeader(r *bufio.Reader) string {
	var header string
	buffer, err := parseSegment(r, patternHeaderLength)
	check(err)

	header = string(buffer)

	return header
}

func parseVersion(r *bufio.Reader) string {

	buffer, err := parseSegment(r, patternVersionLength)
	check(err)

	version := string(buffer)

	return version
}

func parseTracks(r *bufio.Reader, totalLen int32) []track {
	var tracks []track
	var tempTrack track
	var tempTrackLen int32
	for totalLen > 0 {
		tempTrack, tempTrackLen = parseTrack(r)
		tracks = append(tracks, tempTrack)
		totalLen -= tempTrackLen
	}

	fmt.Println(totalLen)

	return tracks
}
