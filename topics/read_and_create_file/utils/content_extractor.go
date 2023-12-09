package utils

import (
	"os"

	"github.com/GuiCintra27/go/topics/read_file/config"
)

func ContentExtractor(sourceFile string) ([]byte, error) {
	logger := config.GetLogger("main")

	//check if the source file exists
	_, err := os.Stat(sourceFile)
	if err != nil {
		logger.Errorf("File '%v' not found", sourceFile)
		return nil , err
	}

	//open and get the content of the source file
	csv, err := os.ReadFile(sourceFile)
	if err != nil{
		logger.Errorf("Error getting content from file: %v", err.Error())
		return nil , err
	}else if len(csv) == 0 {
		logger.Errorf("File '%v' is empty", sourceFile)
		return nil , err
	}

	return csv, nil
}