package piscine

func RetainFirstHalf(str string) string {
	v := len(str)
	char := ""
	if v == 1 {
		return str
	}
	if str == "" {
		return ""
	}
	if v%2 != 0 {
		v = v - 1
	}
	for i := 0; i < v/2; i++ {
		char += string(str[i])
	}
	return char
}

// shorter and more efficient

// package piscine

// func RetainFirstHalf(str string) string {
//     // If it's empty or has 1 character, just return it right away
//     if len(str) <= 1 {
//         return str
//     }
    
//     // Go automatically rounds down integer division! 
//     // e.g., 11 / 2 = 5
//     half := len(str) / 2
    
//     return str[0:half]
// }