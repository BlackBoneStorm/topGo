package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	h := time.Now()
	switch {
	case h.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	fmt.Println(time.Now())

	var n int
	fmt.Scan(&n)
}
