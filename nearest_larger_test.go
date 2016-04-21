package app_academy

import "testing"

func TestNearestLarger(t *testing.T) {
    cases := []struct {
        log string
        arr []int 
        index int
        want int
        err error
    }{
        {"it handles a simple case to the right", []int{2,3,4,8}, 2, 3, nil},
        {"it handles a simple case to the left", []int{2,8,4,3}, 2, 1, nil},
        {"it treats any two larger numbers like a tie", []int{2,6,4,8}, 2, 1, nil},
        {"it chooses the left case in a tie", []int{2,6,4,6}, 2, 1, nil},
        {"it handles a case with an answer > 1 distance to the left", []int{8,2,4,3}, 2, 0, nil},
        {"it handles a case with an answer > 1 distance to the right", []int{2,4,3,8}, 1, 3, nil},
        {"it returns nil if no larger number is found", []int{2,6,4,8}, 3, 0, nil},
    }
    
    for _, c := range cases {
        t.Log(c.log)
        nearest, err := NearestLarger(c.arr, c.index)
        if err != nil {
            if nearest != c.want {
                t.Errorf("nearest equals %d, want %d", nearest, c.want)
            }
        } else {
            if nearest != 0 {
                t.Errorf("nearest equals")
            } 
        }
    }
}