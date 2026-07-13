package main

import (
	"fmt"

)

func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}

func ConcatAlternate(slice1, slice2 []int) []int {
    var first, second []int

    if len(slice2) > len(slice1) {
        first = slice2
        second = slice1
    } else {
        first = slice1
        second = slice2
    }

    result := []int{}
    i := 0

    for i < len(first) && i < len(second) {
        result = append(result, first[i])
        result = append(result, second[i])
        i++
    }

    for i < len(first) {
        result = append(result, first[i])
        i++
    }

    for i < len(second) {
        result = append(result, second[i])
        i++
    }

    return result
}

