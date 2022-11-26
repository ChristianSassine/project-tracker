package db

import (
	"time"

	"github.com/ChristianSassine/projectManager/internal/api"
)

func (db *DB) GetTaskComments(taskId int) (*[]api.Comment, error) {
	comments := []api.Comment{}
	comment := api.Comment{}

	rows, err := db.DB.Query(`
	SELECT users.username, c.content, c.date FROM "comments" c INNER JOIN users ON c.user_id = users.id WHERE c.task_id = $1 ORDER BY c.date ASC`, taskId)

	if err != nil {
		return &comments, err
	}

	for rows.Next() {
		if err := rows.Scan(&comment.Commenter, &comment.Content, &comment.Date); err != nil {
			return &comments, err
		}
		comments = append(comments, comment)
	}

	return &comments, nil
}

func (db *DB) AddTaskComment(userId int, content string, taskId int) error {
	if _, err := db.DB.Exec(`
	INSERT INTO "comments" (user_id, content, date, task_id) VALUES ($1, $2, $3, $4)`, userId, content, time.Now().Local(), taskId); err != nil {
		return err
	}
	return nil
}
