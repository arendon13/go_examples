package main

import (
	"fmt"
	"strconv"
)

// From https://www.hackerrank.com/challenges/time-conversion

func main() {
	// time := "07:05:45PM"
	// time := "12:05:39AM"
	time := "12:45:54PM"

	militaryTime := timeConversion(time)

	fmt.Println(militaryTime)
}

func timeConversion(s string) string {
	hour := s[:2]
	minute := s[3:5]
	second := s[6:8]
	time := s[8:]

	if time == "PM" && hour != "12" {
		hourInt, _ := strconv.Atoi(hour)
		hourInt += 12
		if hourInt == 24 {
			return fmt.Sprintf("%s:%s:%s", "00", minute, second)
		}
		return fmt.Sprintf("%d:%s:%s", hourInt, minute, second)
	} else if time == "AM" && hour == "12" {
		return fmt.Sprintf("%s:%s:%s", "00", minute, second)
	} else {
		return fmt.Sprintf("%s:%s:%s", hour, minute, second)
	}
}
