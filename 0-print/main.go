package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("Hello, we are Holberton School")
	fmt.Printf("the date is %v\n", t)
	fmt.Printf("the year is %v\n", (t.Year()))
	fmt.Printf("the month is %v\n", (t.Month()))
	fmt.Printf("the day is %v\n", (t.Day()))
}
