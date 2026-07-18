package piscine

func FishAndChips(n int) string {
	div2 := n%2 == 0
	div3 := n%3 == 0

	if div2 && div3 {
		return "fish and chips"
	}
	if div2 {
		return "fish"
	}
	if div3 {
		return "chips"
	}
	if n < 0 {
		return "error: number is negative"
	}
	return "error: non divisible"
}
