package app_academy

import "strings"

func LetterCount(str string) map[string]int {
    countedLetters := make(map[string]int)
    for _, s := range strings.Replace(str, " ", "", -1) {
        if _, ok := countedLetters[string(s)]; ok {
            countedLetters[string(s)] = countedLetters[string(s)] + 1
        } else {
            countedLetters[string(s)] = 1
        }
    }
    return countedLetters
}
