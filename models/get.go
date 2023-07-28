package models

import "simpleCrudGo/db"

func Get(id int64) (todos Todo, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM todos WHERE id = $1`, id)

	err = row.Scan(&todos.ID, &todos.Title, &todos.Description, &todos.Done)

	return
}
