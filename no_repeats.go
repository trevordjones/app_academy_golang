package app_academy

import (
    "errors"
    "strconv"
    "strings"
    "fmt"
)

func NoRepeats(beginYear, endYear int) (years []int, err error){
    if beginYear > endYear {
        err = errors.New("End year must be greater than begin year")
    }    
    if err == nil {
        for year := beginYear; year <= endYear; year++ {
            if hasNoDuplicates(year) {
                years = append(years, year)
            }
        }
    }
    if len(years) == 0 {
        years = append(years, 0)
        fmt.Println(years)
    }
    return
}

func hasNoDuplicates(year int) bool {
    yearMap := make(map[string]bool)
    yearStrArray := strings.Split(strconv.Itoa(year), "")
    for _, str := range yearStrArray {
        if yearMap[str] {
            return false
        } else {
            yearMap[str] = true
        }
    }
    return true
}