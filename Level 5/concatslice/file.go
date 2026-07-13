package piscine

func ConcatSlice(slice1, slice2 []int) []int {
	return append(slice1, slice2...)
}
