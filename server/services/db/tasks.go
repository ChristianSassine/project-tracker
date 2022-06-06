package db

import (
	"BugTracker/api"
)

func (db *DB) GetAllTasks(projectId int) (*[]api.Task, error) {
	tasks := []api.Task{}
	task := api.Task{}

	rows, err := db.DB.Query(`
	SELECT id, title, description, creation_date state FROM tasks 
	WHERE project_id = $1`, projectId)

	if err != nil {
		return &[]api.Task{}, err
	}

	for rows.Next() {
		if err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.State, &task.CreationDate); err != nil {
			return &[]api.Task{}, err
		}
		tasks = append(tasks, task)
	}

	return &tasks, nil
}

func (db *DB) GetTasksByState(projectId int, state string) (*[]api.Task, error) {
	tasks := []api.Task{}
	task := api.Task{}

	rows, err := db.DB.Query(`
	SELECT id, title, description, state, creation_date FROM tasks 
	WHERE project_id = $1 AND state = $2
	ORDER BY position ASC`, projectId, state)

	if err != nil {
		return &[]api.Task{}, err
	}

	for rows.Next() {
		if err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.State, &task.CreationDate); err != nil {
			return &[]api.Task{}, err
		}
		tasks = append(tasks, task)
	}

	return &tasks, nil
}

func (db *DB) AddTask(task *api.Task, projectId int) error {
	var position int = 0
	err := db.DB.QueryRow(`SELECT MAX(position) FROM tasks WHERE project_id = $1 AND state = $2`, projectId, task.State).Scan(&position)
	if err != nil {
		// In case it's the first entry, we want the position to be 0
		position = -1
	}
	position++

	_, err = db.DB.Exec(`
	INSERT INTO tasks (title, state, description, creation_date, project_id, position)
	VALUES ($1, $2, $3, $4, $5, $6)`, task.Title, task.State, task.Description, task.CreationDate, projectId, position)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) UpdateTask(task *api.Task, projectId int) error {

	_, err := db.DB.Exec(`
	UPDATE tasks SET title = $1, state = $2, description = $3
	WHERE id = $4 AND project_id = $5`, task.Title, task.State, task.Description, task.Id, projectId)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) UpdateTaskPosition(taskPreviousPosition int, taskCurrentPosition int, taskId int, projectId int) error {
	if taskPreviousPosition < taskCurrentPosition {
		_, err := db.DB.Exec(`
		UPDATE tasks SET position = position - 1 
		WHERE project_id = $1 AND (position > $2) AND (position <= $3)`, projectId, taskPreviousPosition, taskCurrentPosition)
		if err != nil {
			return err
		}
	} else if taskPreviousPosition > taskCurrentPosition {
		_, err := db.DB.Exec(`
		UPDATE tasks SET position = position + 1 
		WHERE project_id = $1 AND (position < $2) AND (position >= $3)`, projectId, taskPreviousPosition, taskCurrentPosition)
		if err != nil {
			return err
		}
	}

	_, err := db.DB.Exec(`
		UPDATE tasks SET position = $1 
		WHERE project_id = $2 AND id = $3`, taskCurrentPosition, projectId, taskId)
	if err != nil {
		return err
	}

	return nil
}

// func (db *DB) UpdateTaskState(taskId int, newState string, taskPreviousPosition int, taskCurrentPosition int, projectId int) error {

// 	_, err := db.DB.Exec(`
// 	UPDATE tasks SET title = $1, state = $2, description = $3
// 	WHERE id = $4 AND project_id = $5`, task.Title, task.State, task.Description, task.Id, projectId)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func (db *DB) DeleteTask(taskId int, projectId int) error {

	position := 0
	err := db.DB.QueryRow(`DELETE FROM tasks WHERE id = $1 AND project_id = $2 RETURNING position`, taskId, projectId).Scan(&position)
	if err != nil {
		return err
	}

	_, err = db.DB.Exec(`
	UPDATE tasks SET position = position - 1 
	WHERE project_id = $1 AND (position > $2)`, projectId, position)
	if err != nil {
		return err
	}

	return nil
}
