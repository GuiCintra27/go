package utils

import (
	"os"
	"strconv"

	"github.com/GuiCintra27/go/topics/read_file/config"
)

func FileCreator(destinationFile string, organizedData *Result)  {
	logger := config.GetLogger("main")

	//create the destination file
	file, err := os.Create(destinationFile)
	if err != nil {
		logger.Errorf("Error creating organized file: %v", err)
		return
	}

	//write the organized data for the destination file
	content := "";
	for i := 0; i < len(organizedData.PersonData); i++ {
		content = content + organizedData.PersonData[i].Name + "," + organizedData.PersonData[i].Age + "," + strconv.Itoa(organizedData.PersonData[i].Score) + "\n"
	}

	//write the content in the destination file
	_, err = file.WriteString(string(organizedData.Title) + "\n" + content)
	if err != nil {
		logger.Errorf("Error filling organized file: %v", err)
		return
	}
}