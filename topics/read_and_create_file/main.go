package main

import (
	"os"

	"github.com/GuiCintra27/go/topics/read_file/utils"
)

func main() {
	// get the source and destination file from the command line
	sourceFile := os.Args[1]
	destinationFile := os.Args[2]

	//get the content  the source file
	csv, err := utils.ContentExtractor(sourceFile)
	if err != nil {
		return
	}

	//organize the data
	organizedData, err := utils.Organizer(string(csv))
	if err != nil {
		return
	}

	// create and fill the destination file
	utils.FileCreator(destinationFile, organizedData)
}