package main

import "math"

func diagonalDifference(arr [][]int32) int32 {
	var n = len(arr)
	var lsum int32 = 0
	var rsum int32 = 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				lsum += arr[i][j]
			}
			if i+j == n-1 {
				rsum += arr[i][j]
			}
		}
	}

	diff := float64(lsum - rsum)
	return int32(math.Abs(diff))
}

func main() {
	a := [][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}
	diagonalDifference(a)
}
