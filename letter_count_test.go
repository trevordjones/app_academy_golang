package app_academy

import "testing"

type Letters map[string]int

func TestLetterCount(t *testing.T) {
    cases := []struct{
        log string
        input string
        expected Letters
    }{
        {"it handles a simple case", "cat", Letters{"c":1, "a":1, "t":1}},
        {"handles multiples of the same letter", "dog", Letters{"d":1, "o":1, "g":1}},
        {"handles a multi-word case", "cats are fun", Letters{"a":2, "c":1, "e":1, "f":1, "n":1, "r":1, "s":1, "t":1, "u":1}},
    }
    for _, c := range cases {
        t.Log(c.log)
        countedLetters := LetterCount(c.input)
        for key, value := range countedLetters {
            if value != c.expected[key] {
                t.Errorf("value is %d, want %d", value, c.expected[key])
            }
        }
    }
}