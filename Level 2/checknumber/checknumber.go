package piscine

func CheckNumber(arg string) bool {
	arr := []rune(arg)
	for i:=0; i < len(arg); i++ {
		if arr[i] > '0' && arr[i] < '9'{
			return true
		}
	}
	return false
}