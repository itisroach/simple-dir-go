package helpers

import (
	"fmt"
	"path"
	"sync"

	"github.com/fatih/color"
)




func Run() {
	// gets the path that user wants to store the folders
	Path := getTextInput("Specify The Path You Want To Store It")
	
	// checks if path that user entered is valid or not
	if !checkIfPathValid(Path) {
		color.Red("Path That You've Entered Does Not Exists")
		return
	}

	// gets user answer if they want to have a parent folder  
	HasParent := getYesNoInput("Do You Want To Create a Parent Folder For It")

	// if user answer was yes
	if HasParent {
		// gets the parent folder name 
		ParentName := getTextInput("Enter The Parent Folder Name")
		// changes path that serveral folder will store in to main_path/parent_folder_name
		Path = path.Join(Path, ParentName)
	}

	// gets number of sub folders user wants to create
	NumberOfFolders := getNumberInput("How Many Folders Do You Want To Create (only numbers allowed)")
	// get the prefix name of folders like Episode 1 => Episode is path prefix
	EachFolderName := getTextInput("Enter Folder Name Prefix (like Episode => which turns to Episode 1)")
	

	var wg sync.WaitGroup

	// iterating over number of folders
	for i := range NumberOfFolders {
		folderName := fmt.Sprintf("%s %d", EachFolderName, i+1)
		folderPath := path.Join(Path, folderName)

		// adding a task for each folder creation
		wg.Add(1)
		go createFolder(folderPath, &wg)
	}
	// waiting till all goroutines ends
	wg.Wait()


	color.Green("Folders Created Successfully")
}