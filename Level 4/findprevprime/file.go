package main

import "fmt"

func main(){
	fmt.Print(FindPrevPrime(7))
}

func IsPrime(num int) bool{
	if num <= 1 {
		return false
	}

	for i := 2 ; i*i <=num ; i++ {
		if num % i == 0 {
			return false
		}
	}
	return true
}

func FindPrevPrime(nb int) int {
	for nb > 1 {
		if IsPrime(nb) {
			return nb
		}
		nb--
	}
	return 0
}
