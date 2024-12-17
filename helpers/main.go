package helpers

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
	"strings"
	"sync"
)



func getTextInput(message string) string {
	var input string

	fmt.Println(message + ":")
	fmt.Scanln(&input)

	if input == "" {
		color.Red("input can not be empty\n\n")
		return getTextInput(message)
	}

	return input
}


func getNumberInput(message string) uint {
	var input uint


	fmt.Println(message + ":")
	fmt.Scanln(&input)

	if input == 0 {
		color.Red("input can not be equal or less than 0\n\n")
		return getNumberInput(message)
	}


	return input
}


// getting user input by choice YES/NO
func getYesNoInput(message string) bool {
	var input string


	fmt.Println(message + "(y/n)")
	fmt.Scanln(&input)


	input = strings.ToLower(input)

	if input == "y"{
		return true
	} else if input == "n" {
		return false
	} else {
		return getYesNoInput(message)
	}
}

// checks if a folder path is valid or a folder exists
func checkIfPathValid(path string) bool {
	if _, err := os.Stat(path); os.IsExist(err) {
		return true
	}

	return false
}


func createFolder(path string, wg *sync.WaitGroup) {
	// checking if folder does not exists
	if checkIfPathValid(path) {
		return
	}

	// creating folders
	err := os.MkdirAll(path, 0777)

	color.Green("creating " + path + " folder ...")

	// checking if folder created successfully
	if err != nil {
		log.Fatal(err)
	}

	// makes a wait group done after function ends 
	defer wg.Done()
}