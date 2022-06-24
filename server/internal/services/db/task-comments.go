package db

import (
	"time"

	"github.com/krispier/projectManager/internal/api"
)

func (db *DB) GetTaskComments(taskId int) (*[]api.Comment, error) {
	comments := []api.Comment{}
	comment := api.Comment{}

	rows, err := db.DB.Query(`
	SELECT users.username, comments.content, comments.date FROM comments
	INNER JOIN users ON comments.user_id = users.id
	WHERE task_id = $1 ORDER BY comments.date ASC`, taskId)

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
	INSERT INTO comments (user_id, content, date, task_id)
	VALUES ($1, $2, $3, $4)`, userId, content, time.Now().Local(), taskId); err != nil {
		return err
	}
	return nil
}
