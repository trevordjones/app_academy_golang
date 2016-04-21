package app_academy

import "testing"

func TestWordUnscrambler(t *testing.T) {
    cases := []struct{
        log string
        word string
        dictionary []string
        expected []string
    }{
        {"it handles a simple word", "cat", []string{"tac"}, []string{"tac"}},
        {"it does not inappropriately choose a word", "cat", []string{"tom"}, []string{""}},
        {"it handles a simple word from a larger dictionary", "cat", []string{"tic", "toc", "tac", "toe"}, []string{"tac"}},
        {"it only handles words of the correct length", "cat", []string{"scatter", "tac", "ca"}, []string{"tac"}},
        {"it handles multiple successful cases", "turn", []string{"numb", "turn", "runt", "nurt"}, []string{"turn", "runt", "nurt"}},
    }
    
    for _, c := range cases {
        t.Log(c.log)
        unscrambledWords := WordUnscrambler(c.word, c.dictionary)
        for i, value := range c.expected {
            if unscrambledWords[i] != value {
                t.Errorf("unscrambled word is %q, expected %q", unscrambledWords[i], value)
            }
        }
    }
}    
