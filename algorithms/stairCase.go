package main

import "fmt"

func staircase(n int32) {
	for i := 0; i < int(n); i++ {
		for j := i; j < int(n)-1; j++ {
			fmt.Printf(" ")
		}
		for j := int(n); j >= int(n)-i; j-- {
			fmt.Printf("#")
		}

		fmt.Println()
	}
}

func main() {
	staircase(4)
}
