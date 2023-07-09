package main

import "fmt"

func main() {
	fmt.Println(euclid(12, 9))
	fmt.Println(rev("oopa"))
}

/*find the GCD*/
func euclid(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func rev(word string) string {
	var ans string
	for i := len(word) - 1; i >= 0; i-- {
		ans += string(word[i])
	}
	return ans
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
