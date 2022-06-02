package main

import "fmt"

func plusMinus(arr []int32) {
	var positive, negative, zero float64 //variable
	element := float64(len(arr))         //element size

	for _, i := range arr { //for loop
		if i > 0 {
			positive++
		} else if i < 0 {
			negative++
		} else {
			zero++
		}
	}

	//6 decimal number
	fmt.Printf("%.6f\n", positive/element)
	fmt.Printf("%.6f\n", negative/element)
	fmt.Printf("%.6f\n", zero/element)

}

func main() {
	arr := []int32{1, 1, 0, -1, -1}
	plusMinus(arr)
}
