package app_academy

import "errors"

func NearestLarger(intArray []int, index int) (nearest int, err error) {
    for i, value := range intArray {
        if value > intArray[index] {
            nearest = i
            err = errors.New("No error")
            break 
        }
    }
    return 
}