package piscine

func CanJump(arr []uint) bool {
	// Empty array
	if len(arr) == 0 {
		return false
	}

	// Single element = already at the end
	if len(arr) == 1 {
		return true
	}

	pos := 0
	last := len(arr) - 1

	for {
		// If we are already at the last index
		if pos == last {
			return true
		}

		// Jump exactly arr[pos] steps
		next := pos + int(arr[pos])

		// If jump goes out of bounds
		if next < 0 || next >= len(arr) {
			return false
		}

		pos = next
	}
}