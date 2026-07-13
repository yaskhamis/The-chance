package main

func main() {
	
}

func SaveAndMiss(arg string, num int) string {
	// If num is 0 or negative, return original string
	if num <= 0 {
		return arg
	}

	result := ""
	save := true

	for i := 0; i < len(arg); i += num {
		end := i + num
		if end > len(arg) {
			end = len(arg)
		}

		if save {
			result += arg[i:end]
		}

		// Toggle between save and miss
		save = !save
	}

	return result
}