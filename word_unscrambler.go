package app_academy

import (
    "strings"
    "sort"
)

func WordUnscrambler(word string, dictionary []string) (unscrambledWord []string){
    for _, str := range dictionary {
        if sortString(word) == sortString(str) {
            unscrambledWord = append(unscrambledWord, str)
        }
    }
    if unscrambledWord == nil {
        unscrambledWord = append(unscrambledWord, "")
    }
    return
}

func sortString(str string) string{
    s := strings.Split(str, "")
    sort.Strings(s)
    return strings.Join(s, "")
}