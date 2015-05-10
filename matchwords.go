package natlang

import "regexp"

var WordRegexp = regexp.MustCompile(`\pL+('\pL+)*`)

// MatchWords matches the words in a text, including any embedded apostrophes.
func MatchWords(text string) []string {
	return WordRegexp.FindAllString(text, -1)
}
