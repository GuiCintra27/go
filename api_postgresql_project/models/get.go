package models

import "github.com/GuiCintra27/go/api_postgresql_project/db"

func Get(id int64) (todo Todo, err error) {
	conn, err := db.Connect()

	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow("SELECT * FROM todos WHERE id = $1", id)
	
	err = row.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Done)

	return

}