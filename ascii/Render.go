package ascii

func RenderLine(text string, banner map[rune][]string) []string {
	word := make([]string, 8)
	for i := 0; i <= 7; i++ {
		for _, ch := range text {
			word[i] += banner[ch][i]
		}
	}
	return word
}
