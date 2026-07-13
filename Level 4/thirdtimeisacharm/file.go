package main

func ThirdTimeIsACharm(str string) string {
	result := ""

	for i := 2; i < len(str); i += 3 {
		result += string(str[i])
	}

	if result == "" {
		return "\n"
	}

	return result + "\n"
}