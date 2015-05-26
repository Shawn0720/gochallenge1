package drum

import (
	"bytes"
	"fmt"
)

func printSteps(steps [16]bool) string {
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

func printVersion(version string) string {
	return fmt.Sprintf("Saved with HW Version: %s\n", version)
}

func printTempo(tempo float32) string {
	return fmt.Sprintf("Tempo: %g\n", tempo)
}

func printTrack(trackObj track) string {
	return fmt.Sprintf("(%d) %s\t%s\n", trackObj.id, trackObj.name, printSteps(trackObj.steps))
}
