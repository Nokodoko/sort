package main

import (
	"fmt"
)

func main() {
	add(5, 7)
	sum(15)
}

/*O(1) - constant time*/
func add(a, b int) int {
	fmt.Println(add)
	return a + b
}

/*O(N) - N will be the determining factor to how long this will take*/
func sumToMax(max int) int {
	sum := 0
	for i := 0; i >= max; i++ {
		sum += i
	}
	fmt.Println(sum)
	return sum
}

/*O(1) - constant time (less calculation than the for loop above)*/
func sumToMax2(max int) int {
	return (max * (max + 1)) / 2
}

/* O(N) - value of the inopute dictates how fast the code will run -  N == len(vals) Linear Time*/
func sumVals(vals []int) int {
	var sum int
	for _, v := range vals {
		sum += val
	}
	return sum
}

func Find(list []int, x int) int {
	for i, v := range list {
		switch {
		case i == x:
			fmt.Println("Matched", x)
			return i
		default:
			fmt.Println("There is no mathcing", x, "here.")
		}
	}
	return -1
}

func Grid(x, y int) string {

	return ""
}
