package main

func simpleArraySum(arr []int32) int32 {
	var sum int32 = 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func main() {
	a := []int32{2, 3, 5, 11}
	simpleArraySum(a)
}
