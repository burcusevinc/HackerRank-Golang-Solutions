package main

func compareTriplets(a []int32, b []int32) []int32 {

	var alice int32 = 0
	var bob int32 = 0
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			alice++
		} else if a[i] < b[i] {
			bob++
		}
	}
	return []int32{alice, bob}

}

func main() {
	a := []int32{2, 3, 5}
	b := []int32{2, 1, 13}
	compareTriplets(a, b)
}
