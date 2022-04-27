package map_demo

import "testing"

func TestXxx(t *testing.T) {
	s := "anagram"
	tx := "nagaram"
	if !isAnagram(s, tx) {
		t.Errorf("%s is anagram of %s", s, tx)
	}
	s = "rat"
	tx = "car"
	if isAnagram(s, tx) {
		t.Errorf("%s is not anagram of %s", s, tx)
	}
}
