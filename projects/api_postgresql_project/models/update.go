package models

import "github.com/GuiCintra27/go/projects/api_postgresql_project/db"

func Update(id int64, todo Todo) (int64, error) {
	conn, err := db.Connect()

	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec("UPDATE todos SET done = $1 WHERE id = $2", todo.Done, id)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}