package usecase

import (
	"database/sql"
	"github.com/alaminmahamud/todo-app/entity"
)

func GetTasks(db *sql.DB) entity.TaskCollection {
	sql       := "SELECT * FROM tasks;"
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()
	result := entity.TaskCollection{}

	for rows.Next() {
		task := entity.Task{}
		err  := rows.Scan(
			&task.ID,
			&task.Name,
			&task.Done,
			)

		if err != nil {
			panic(err)
		}

		result.Tasks = append(result.Tasks, task)
	}

	return result
}