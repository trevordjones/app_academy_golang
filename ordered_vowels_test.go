package app_academy

import "testing"

func TestOrderedVowels(t *testing.T) {
    cases := []struct{
        log string
        input string
        expected string
    }{
        {"it returns a word that is in order", "amends", "amends"},
        {"it does not return a word that is not in order", "complicated", ""},
        {"handle double vowels", "afoot", "afoot"},
        {"handles a word with a single vowel", "ham", "ham"},
        {"handles a word with no vowel", "crypt", "crypt"},
        {"handles a word with a single letter", "o", "o"},
        {"ignores the letter y", "tamely", "tamely"},
        {"processes a string with several words", "this is a test of the vowel ordering system", "this is a test of the system"},
    }
    for _, c := range cases {
        t.Log(c.log)
        orderedVowels := OrderedVowels(c.input)
        if orderedVowels != c.expected {
            t.Errorf("orderedVowels equals %s, expected %s", orderedVowels, c.expected)
        }
    }
}