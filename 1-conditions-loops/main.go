//prints some facts about the school and founders

package main

import (
	"fmt"
	"math/rand"
)

func main() {

	randomNumber := rand.Intn(100)
	if randomNumber > 50 {
		fmt.Printf("my random number is %d and is greater than 50\n", randomNumber)
	} else {
		fmt.Printf("my random number is %d and is less than 50\n", randomNumber)
	}

	school := "Holberton School"
	if school == "Holberton School" {
		fmt.Printf("I am a student of the %s\n", school)
	} else {
		fmt.Printf("I am a student at Hogwarts!\n")
	}

	beautifulWeather := true
	if beautifulWeather == true {
		fmt.Printf("It's a beautiful weather!\n")
	} else {
		fmt.Printf("It's a disappointing weather!\n")
	}

	holbertonFounders := [3]string{"Rudy Rigot", "Sylvain Kalache", "Julien Barbier"}
	i := 0
	for ; i < 3; i++ {
		fmt.Printf("%s is a founder\n", holbertonFounders[i])
	}
	// fmt.Printf("boolean value sentence\n")
	// fmt.Printf("is a founder\n")
	// fmt.Printf("is a founder\n")
	// fmt.Printf("is a founder\n")
}
