package helpers

import (
	"fmt"
	"strings"
)



func getTextInput(message string) string {
	var input string

	fmt.Println(message)
	fmt.Scanln(&input)

	if input == "" {
		fmt.Println("input can not be empty")
		getTextInput(message)
	}

	return input
}


func getNumberInput(message string) uint {
	var input uint

	if input == 0 {
		fmt.Println("input can not be equal or less than 0")
		getNumberInput(message)
	}

	fmt.Println(message)
	fmt.Scanln(&input)

	return input
}


func getYesNoInput(message string) bool {
	var input string


	fmt.Println(message)
	fmt.Scanln(&input)


	input = strings.ToLower(input)

	if input == "y"{
		return true
	} else if input == "n" {
		return false
	} else {
		val := getYesNoInput(message)
		return val
	}
}
