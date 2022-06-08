package main

func birthdayCakeCandles(candles []int32) int32 {

	var max int32 = 0
	var count int32 = 0

	for i := 0; i < len(candles); i++ {
		if candles[i] > max {
			max = candles[i]
		}
	}

	for i := 0; i < len(candles); i++ {
		if candles[i] == max {
			count++
		}
	}

	return count

}

func main() {
	a := []int32{4, 4, 1, 3}
	birthdayCakeCandles(a)
}
