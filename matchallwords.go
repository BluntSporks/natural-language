package natlang

import "regexp"

var AllWordRegexp = regexp.MustCompile(`\pL+(\pP+\pL+)*`)

// MatchAllWords matches the words in a text, including any embedded apostrophes.
// It also matches wordlike expressions that contain nonword characters, such as email or web addresses.
func MatchAllWords(text string) []string {
	return AllWordRegexp.FindAllString(text, -1)
}
