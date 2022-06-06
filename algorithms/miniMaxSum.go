package main

import (
	"fmt"
	"sort"
)

func miniMaxSum(arr []int) {

	sort.Ints(arr)

	var sum int
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	fmt.Println(sum-arr[len(arr)-1], sum-arr[0])

	//totalMin := arr[0] + arr[1] + arr[2] + arr[3]
	//totalMax := arr[1] + arr[2] + arr[3] + arr[4]
	//fmt.Println(totalMin, totalMax)

}

func main() {
	a := []int{1, 3, 5, 7, 9}
	miniMaxSum(a)
}
