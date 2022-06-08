package main

import (
	"fmt"
	"strconv"
	"strings"
)

func timeConversion(s string) string {

	pm := strings.HasSuffix(s, "PM")
	am := strings.HasSuffix(s, "AM")

	// am and pm part of the string removed
	time := s[:len(s)-2]

	// splits string when seeing ":"
	timeArr := strings.Split(time, ":")
	hour, minute, second := timeArr[0], timeArr[1], timeArr[2]

	hourInt, _ := strconv.Atoi(hour) // string to int

	if am && hourInt == 12 {
		hourInt = 0
	}

	if pm && hourInt != 12 {
		hourInt += 12
	}

	hour = strconv.Itoa(hourInt) // int to string

	return fmt.Sprintf("%02s:%02s:%02s", hour, minute, second)
}

func main() {
	timeConversion("07:05:45PM")
}
