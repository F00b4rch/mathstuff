package main

import (
	"flag"
	"fmt"
)

func main() {
	sum := 0

	flagNum := flag.Int("number", 42, "your starting number")
	flag.Parse()

	num := *flagNum

	// Stop exercise if number 4 is reached
	limit := 1

	fmt.Println("Starting loop")
	for {
		sum++
		// init
		isEven := evenNumber(num)
		fmt.Printf("turn %v, number is %v\n", sum, num)

		if isEven == true {
			num = even(num)
			if num == limit {
				fmt.Printf("Ending here with final value %v, after %v turns\n", num, sum)
				break
			}
		} else if isEven == false {
			num = nEven(num)
			if num == limit {
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
	//fmt.Printf("%v is new value\n", res)
	return res
}

// Not even number
func nEven(n int) int {
	res := n*3 + 1
	//fmt.Printf("%v is new value\n", res)
	return res
}
