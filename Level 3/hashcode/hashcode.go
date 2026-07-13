package piscine

func HashCode(dec string) string {
	arr := []rune(dec)
	result := ""
	for _, v := range arr {
		temp := int(v) + len(arr)%127
		if temp <= 32 {
			temp += 33
		}
		result += string(temp)
	}
	return result
}
