package ascii

import (
	"strings"
)

func GenerateArt(input string, banner map[rune][]string) string {
	if input == "" {
		return ""
	}

	for _, rn := range input {
		if (rn < 32 || rn > 126) && rn != '\r' && rn != '\n' {
			return "Invalid Ascii character"
		}
	}

	line := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
	var result strings.Builder

	for _, st := range line {
		if st == "" {
			result.WriteString("\n")
		} else {
			rend := RenderLine(st, banner)
			for _, l := range rend {
				result.WriteString(l)
				result.WriteString("\n")
			}
		}
	}
	return result.String()
}
