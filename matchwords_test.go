package natlang

import "testing"

func TestMatchWords(t *testing.T) {
	text := "Don't e-mail me your Ph.D. to user@example.com, please!"
	exp := []string{"Don't", "e", "mail", "me", "your", "Ph", "D", "to", "user", "example", "com", "please"}
	act := MatchWords(text)
	if len(exp) != len(act) {
		t.Error("Expected length", len(exp), "got", len(act))
		return
	}
	for i := 0; i < len(exp); i++ {
		if exp[i] != act[i] {
			t.Error("Expected", exp[i], "got", act[i])
		}
	}
}
