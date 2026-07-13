package piscine

func FirstWord(s string) string {
	arr := []rune(s)
	char := ""
	for i := 0; i < len(arr); i++ {
		if arr[i] != ' ' {
			char += string(arr[i])
		} else {
			break
		}
	}
	
	return char + "\n"
}
