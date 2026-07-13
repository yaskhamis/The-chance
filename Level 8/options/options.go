package group

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
		return
	}

	var opts int32 = 0

	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]

		if len(arg) < 2 || arg[0] != '-' {
			fmt.Println("Invalid Option")
			return
		}

		// -h flag has highest priority if first
		if arg == "-h" {
			fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
			return
		}

		// Process each character in the option
		for j := 1; j < len(arg); j++ {
			ch := arg[j]

			if ch < 'a' || ch > 'z' {
				fmt.Println("Invalid Option")
				return
			}

			// Set the bit: a = 0, b = 1, ..., z = 25
			bit := ch - 'a'
			opts |= 1 << bit
		}
	}

	// Convert int32 to 32-bit binary string
	bits := fmt.Sprintf("%032b", opts)

	// Split into 4 groups of 8
	fmt.Printf("%s %s %s %s\n", bits[:8], bits[8:16], bits[16:24], bits[24:])
}