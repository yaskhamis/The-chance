package group

func ItoaBase(value, base int) string {
	const digits = "0123456789ABCDEF"

	// Special case for zero
	if value == 0 {
		return "0"
	}

	result := ""
	negative := false

	// Handle negative values
	if value < 0 {
		negative = true
		value = -value
	}

	// Convert number to base
	for value > 0 {
		remainder := value % base
		result = string(digits[remainder]) + result
		value /= base
	}

	// Add minus sign if needed
	if negative {
		result = "-" + result
	}

	return result
}
