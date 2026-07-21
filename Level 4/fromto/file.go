package piscine

import "strconv"

func FromTo(from int, to int) string {
	if from > 99 || from < 0 || to > 99 || to < 0 {
		return "Invalid\n"
	}
	step := 1
	if from > to {
		step = -1
	}
	res := ""
	for i := from; ; i += step {
		s := strconv.Itoa(i)
		if i < 10 {
			s = "0" + s
		}
		if res != "" {
			res += ", "
		}
		res += s
		if i == to {
			break
		}
	}
	return res + "\n"
}
