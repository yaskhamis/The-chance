package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(2, 1))
	// fmt.Print(FromTo(2, 2))
	// fmt.Print(FromTo(100, 10))
}

func FromTo(from int, to int) string {
	if from > 99 || from < 0 || to > 99 || to < 0 {
		return "Invalid" + "\n"
	}
	result := ""
	if from > to {
		for i := from ; i >= to ; i-- {
			if i < 10 {
				if i != to {
					result += "0" + strconv.Itoa(i) + ", "
				} else {
					result += "0" + strconv.Itoa(i) + "\n"
				} 
			} else if i >= 10 {
				if i != to {
					result += strconv.Itoa(i) + ", "
				} else {
					result += strconv.Itoa(i) + "\n"
				}
			}
		}
	} else if from < to {

	} else  {

	}
	return result
}

