package main

import "fmt"

func gradingStudents(grades []int32) []int32 {
	var finalGrades []int32 //empty array

	for i := 0; i < len(grades); i++ {
		if grades[i] >= 38 && grades[i]%5 >= 3 {
			grades[i] = grades[i] + 5 - grades[i]%5 // round our number
		}
		finalGrades = append(finalGrades, grades[i])
	}
	fmt.Println(finalGrades)
	return finalGrades

}

func main() {
	a := []int32{4, 73, 67, 38, 33}
	gradingStudents(a)
}
