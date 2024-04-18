package utils

import (
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/GuiCintra27/go/topics/read_and_create_file/config"
)

type Result struct {
	Title string
	PersonData []PersonData
}

type PersonData struct {
	Name string
	Age  string
	Score int
}

func Organizer(str string) (*Result, error) {
	logger := config.GetLogger("main")
	data := []PersonData{}

	var title string

	p, _ := regexp.Compile("([A-z]+),([0-9]+),([0-9]+)")
	t, _ := regexp.Compile("([A-z]+),([A-z]+),([A-zÀ-ú]+)")

	pStr := p.FindAllString(str, -1)
	title = t.FindString(str)

	for i := 0; i < len(pStr); i++ {
		splited := strings.Split(pStr[i], ",")

		score, _ := strconv.Atoi(splited[2])

		newPersonData := &PersonData{
			Name: splited[0],
			Age:   splited[1],
			Score: score,				
		}
		data = append(data, *newPersonData)
	}

	sort.SliceStable(data, func(i, j int) bool {
		if data[i].Name == data[j].Name {
			if data[i].Age == data[j].Age {
				return data[i].Score < data[j].Score
			}
			return data[i].Age < data[j].Age
		}
		return data[i].Name < data[j].Name
	})

	if title == "" {
		logger.Error("Error organizing file content: title empty or don't follow the pattern")
		print("Title pattern: Nome,Idade,Pontuação")
		
	}

	return &Result{
		Title: title,
		PersonData: data,
	}, nil
}