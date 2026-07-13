package piscine

func LastWord(s string) string {
	arr := []rune(s)
	char := ""
	for i := len(arr)-1; i >= 0 ; i-- {
		if arr[i] != ' ' || arr[len(arr)-1] == ' '{
			char += string(arr[i])
		} else {
			break
		}
	}
	
	fil := []rune(char)
	jh := ""
	for j := len(fil)-1; j >= 0 ; j-- {
		if fil[j] != ' '{
			jh += string(fil[j])
		}
	}
	
	return jh + "\n"
}
