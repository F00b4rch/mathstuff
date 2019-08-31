package main

import "fmt"

func main() {
	sum := 0
	num := 0

	fmt.Println("Starting loop")
	for {
		sum++
		// init
		if num == 0 {
			num = 13
		}
		isEven := evenNumber(num)
		fmt.Printf("turn %v, number is %v\n", sum, num)

		if isEven == true {
			num = even(num)
		} else if isEven == false {
			num = nEven(num)
			if num == 4 {
				fmt.Printf("Ending here with final value %v, after %v turns\n", num, sum)
				break
			}
		}
	}
}

// Testing even number
func evenNumber(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

// Even number
func even(n int) int {
	res := n / 2
	fmt.Printf("%v is new value\n", res)
	return res
}

// Not even number
func nEven(n int) int {
	res := n*3 + 1
	fmt.Printf("%v is new value\n", res)
	return res
}
