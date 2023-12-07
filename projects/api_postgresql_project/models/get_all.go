package models

import (
	"github.com/GuiCintra27/go/projects/api_postgresql_project/db"
)


func GetAll() (todos []Todo, err error) {
	conn, err := db.Connect()
	
	if err != nil {
		return
	}
	defer conn.Close()

	row, err := conn.Query("SELECT * FROM todos")

	if err != nil {
		return
	}

	for row.Next() {
		var todo Todo

		err = row.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Done)

		if err != nil {
			continue
		}

		todos = append(todos, todo)
	}

	return

}