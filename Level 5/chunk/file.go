package piscine

import "fmt"

func Chunk(slice []int, size int) {
	if size <= 0 {
		fmt.Println()
		return
	}

	res := [][]int{}

	for i := 0; i < len(slice); i += size {
		end := i + size

		if end > len(slice) {
			end = len(slice)
		}
		res = append(res, slice[i:end])
	}
	fmt.Println(res)
}
