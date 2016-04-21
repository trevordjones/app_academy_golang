package app_academy

import "testing"

func TestNoRepeats(t *testing.T) {
    cases := []struct{
        log string
        beginYear int
        endYear int
        expected []int
    } {
        {"it returns a no repeat year", 1234, 1234, []int{1234}},
        {"it does not return a repeat year", 1123, 1123, []int{0}},
        {"it returns only those years that have no repeated digits", 1980, 1989, []int{1980, 1982, 1983, 1984, 1985, 1986, 1987}},
    }
    
    for _, c := range cases {
        t.Log(c.log)
        noRepeats, err := NoRepeats(c.beginYear, c.endYear)
        for i := 0; i < len(c.expected); i++ {
            if err == nil {
                if noRepeats[i] != c.expected[i] {
                    t.Errorf("noRepeats equals %d, want %d", noRepeats[i], c.expected[i])
                }
            }
        }
    }
}