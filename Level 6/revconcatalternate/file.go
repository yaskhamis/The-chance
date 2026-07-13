package main

import (
	"fmt"

)

func main() {
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))
}

func RevConcatAlternate(slice1, slice2 []int) []int {
	len1 := len(slice1)
	len2 := len(slice2)
	i1 := len1 - 1
	i2 := len2 - 1
	result := []int{}

	for i1 >= 0 || i2 >= 0 {
		rem1 := i1 + 1
		rem2 := i2 + 1

		if rem1 > rem2 && i1 >= 0 { // slice1 is bigger
			result = append(result, slice1[i1])
			i1--
		} else if rem2 > rem1 && i2 >= 0 { // slice2 is bigger
			result = append(result, slice2[i2])
			i2--
		} else if rem1 == rem2 { // remaining lengths equal → alternate
			if i1 >= 0 {
				result = append(result, slice1[i1])
				i1--
			}
			if i2 >= 0 {
				result = append(result, slice2[i2])
				i2--
			}
		}
	}
	return result
}