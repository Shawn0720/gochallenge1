package drum

import (
	"bytes"
	"fmt"
)

func print_steps(steps [16]bool) string {
	var buffer bytes.Buffer

	buffer.WriteString("|")
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if steps[i*4+j] {
				buffer.WriteString("x")
			} else {
				buffer.WriteString("-")
			}
		}
		buffer.WriteString("|")
	}
	return buffer.String()
}

func print_version(version string) string {
	return fmt.Sprintf("Saved with HW Version: %s\n", version)
}

func print_tempo(tempo float32) string {
	return fmt.Sprintf("Tempo: %g\n", tempo)
}

func print_track(track Track) string {
	return fmt.Sprintf("(%d) %s\t%s\n", track.id, track.name, print_steps(track.steps))
}
