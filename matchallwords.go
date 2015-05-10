package natlang

import "regexp"

// MatchAllWords matches the words in a text, including any embedded apostrophes.
// It also matches wordlike expressions that contain nonword characters, such as email or web addresses.
func MatchAllWords(text string) []string {
	re := regexp.MustCompile(`\pL+(\pP+\pL+)*`)
	return re.FindAllString(text, -1)
}
