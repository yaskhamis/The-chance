package piscine

func RectPerimeter(w, h int) int {
	if w < 0 || h < 0{
		return -1
	}
	n:= 2*(w+h)
	return n
}