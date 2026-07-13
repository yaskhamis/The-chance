package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Invalid input.")
		return
	}

	arrStr := os.Args[1]
	targetStr := os.Args[2]

	// ---------- Validate array format ----------
	if len(arrStr) < 2 || arrStr[0] != '[' || arrStr[len(arrStr)-1] != ']' {
		fmt.Println("Invalid input.")
		return
	}

	// ---------- Parse target ----------
	if strings.Contains(targetStr, " ") {
		fmt.Println("Invalid target sum.")
		return
	}

	target, err := strconv.Atoi(targetStr)
	if err != nil {
		fmt.Println("Invalid target sum.")
		return
	}

	// ---------- Parse array ----------
	content := strings.TrimSpace(arrStr[1 : len(arrStr)-1])
	if content == "" {
		fmt.Println("No pairs found.")
		return
	}

	parts := strings.Split(content, ",")
	arr := []int{}

	for _, p := range parts {
		p = strings.TrimSpace(p)
		n, err := strconv.Atoi(p)
		if err != nil {
			fmt.Println("Invalid number:", p)
			return
		}
		arr = append(arr, n)
	}

	// ---------- Find pairs ----------
	pairs := [][]int{}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				pairs = append(pairs, []int{i, j})
			}
		}
	}

	// ---------- Output ----------
	if len(pairs) == 0 {
		fmt.Println("No pairs found.")
		return
	}

	fmt.Printf("Pairs with sum %d: %v\n", target, pairs)
}
