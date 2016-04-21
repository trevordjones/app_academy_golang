package app_academy

import (
    "strings"
    "sort"
)

func OrderedVowels(str string) string {
    vowels := map[string]bool{"a":true, "e":true, "i":true, "o":true, "u":true}
    var orderedArray []string
    var vowelsFromWord []string
    for _, word := range strings.Split(str, " ") {
        vowelsFromWord = nil
        for _, w := range word {
            if _, ok := vowels[string(w)]; ok {
                vowelsFromWord = append(vowelsFromWord, string(w))
            }
        }
        if sort.StringsAreSorted(vowelsFromWord) {
            orderedArray = append(orderedArray, word)
        }
    }
    return strings.Join(orderedArray, " ")
}