package piscine

func DigitLen(n, base int) int {
	if !(base >= 2 && base <= 36) {
		return -1
	}
	if n < 0 {
		n *= -1
	}
	x := 0

	for n != 0 {
		n /= base
		x++
	}
	return x
}