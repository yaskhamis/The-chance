package piscine

func CountChar(str string, c rune) int {
	arr := []rune(str)
	count := 0
	if str == "" {
		return 0
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] == c {
			count++
		}
	}
	return count
}
