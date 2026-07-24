package piscine

func CanJump(arr []uint) bool {
	if len(arr) == 0 {
		return false
	}
	if len(arr) == 1 {
		return true
	}

	pos := 0
	last := len(arr) - 1

	for i := 0; i <= last && i <= pos; i++ {
		next := i + int(arr[i])

		if next > pos {
			pos = next
		}
		if pos >= last {
			return true
		}
	}
	return false
}
