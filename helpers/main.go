package helpers

import "fmt"



func getUserInput(message string) string {
	var input string

	fmt.Println(message)
	fmt.Scanln(&input)


	return input
}