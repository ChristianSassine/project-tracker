package db

import (
	"time"

	"github.com/krispier/projectManager/internal/api"

	"github.com/lib/pq"
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

func (db *DB) GetTasksWithLimit(projectId int, limit int) (*[]api.Task, error) {
	tasks := []api.Task{}
	task := api.Task{}

	rows, err := db.DB.Query(`
	SELECT id, title, description, state, creation_date FROM tasks 
	WHERE project_id = $1
	ORDER BY creation_date DESC
	LIMIT $2`, projectId, limit)

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

func (db *DB) GetTasksStats(projectId int) (*api.TaskStats, error) {
	taskStats := api.TaskStats{}
	statsArray := []int64{0, 0, 0}

	if err := db.DB.QueryRow(`
	SELECT(
		ARRAY[(SELECT COUNT(*) FROM tasks WHERE state = 'TODO' AND project_id = $1),
		(SELECT COUNT(*) FROM tasks WHERE state = 'ONGOING' AND project_id = $1),
		(SELECT COUNT(*) FROM tasks WHERE state = 'DONE' AND project_id = $1)]
	)`, projectId).Scan(pq.Array(&statsArray)); err != nil {
		return &api.TaskStats{}, err
	}

	taskStats.TodoTasks = statsArray[0]
	taskStats.OngoingTasks = statsArray[1]
	taskStats.DoneTasks = statsArray[2]

	return &taskStats, nil
}

func (db *DB) AddTask(task *api.Task, projectId int) error {
	var position int = 0
	if err := db.DB.QueryRow(`SELECT MAX(position) FROM tasks WHERE project_id = $1 AND state = $2`, projectId, task.State).Scan(&position); err != nil {
		// In case it's the first entry, we want the position to be 0
		position = -1
	}
	position++

	_, err := db.DB.Exec(`
	INSERT INTO tasks (title, state, description, creation_date, project_id, position)
	VALUES ($1, $2, $3, $4, $5, $6)`, task.Title, task.State, task.Description, time.Now().Local(), projectId, position)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) UpdateTask(task *api.Task, projectId int) error {

	_, err := db.DB.Exec(`
	UPDATE tasks SET title = $1, description = $2
	WHERE id = $3 AND project_id = $4`, task.Title, task.Description, task.Id, projectId)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) UpdateTaskPosition(taskPreviousPosition int, taskCurrentPosition int, taskId int) error {

	var currentState string
	var projectId int
	if err := db.DB.QueryRow(`
	SELECT state, project_id
	FROM tasks WHERE id = $1
	`, taskId).Scan(&currentState, &projectId); err != nil {
		return err
	}

	// Incrementing or decrementing the position of the other elements of the array depending on the positions
	if taskPreviousPosition < taskCurrentPosition {
		_, err := db.DB.Exec(`
		UPDATE tasks SET position = position - 1 
		WHERE project_id = $1 AND state = $2 AND (position > $3) AND (position <= $4)`, projectId, currentState, taskPreviousPosition, taskCurrentPosition)
		if err != nil {
			return err
		}
	} else if taskPreviousPosition > taskCurrentPosition {
		_, err := db.DB.Exec(`
		UPDATE tasks SET position = position + 1 
		WHERE project_id = $1 AND state = $2 AND (position < $3) AND (position >= $4)`, projectId, currentState, taskPreviousPosition, taskCurrentPosition)
		if err != nil {
			return err
		}
	}

	// Setting the task's current position
	_, err := db.DB.Exec(`
		UPDATE tasks SET position = $1 
		WHERE project_id = $2 AND id = $3`, taskCurrentPosition, projectId, taskId)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) UpdateTaskState(newState string, taskCurrentPosition int, taskId int, projectId int) (*api.Task, error) {

	// Decrementing the position for the array where it left
	previousTask := &api.Task{}
	previousPosition := 0

	if err := db.DB.QueryRow(`
	SELECT state, title, position 
	FROM tasks WHERE project_id = $1 AND id = $2
	`, projectId, taskId).Scan(&previousTask.State, &previousTask.Title, &previousPosition); err != nil {
		return &api.Task{}, err
	}

	_, err := db.DB.Exec(`
	UPDATE tasks SET position = position - 1
	WHERE project_id = $1 AND state = $2 AND (position > $3)`, projectId, previousTask.State, previousPosition)
	if err != nil {
		return &api.Task{}, err
	}

	// Incrementing the position for the array where it's added
	_, err = db.DB.Exec(`
		UPDATE tasks SET position = position + 1 
		WHERE project_id = $1 AND (position >= $2) AND state = $3`, projectId, taskCurrentPosition, newState)
	if err != nil {
		return &api.Task{}, err
	}

	// Setting the new values
	_, err = db.DB.Exec(`
	UPDATE tasks SET state = $1, position = $2
	WHERE id = $3 AND project_id = $4`, newState, taskCurrentPosition, taskId, projectId)
	if err != nil {
		return &api.Task{}, err
	}

	return previousTask, nil
}

func (db *DB) DeleteTask(taskId int, projectId int) (string, error) {

	position := 0
	title := ""
	state := ""

	err := db.DB.QueryRow(`DELETE FROM tasks WHERE id = $1 AND project_id = $2 RETURNING position, title, state`, taskId, projectId).Scan(&position, &title, &state)
	if err != nil {
		return "", err
	}

	_, err = db.DB.Exec(`
	UPDATE tasks SET position = position - 1 
	WHERE project_id = $1 AND state = $2 AND (position > $3)`, projectId, state, position)
	if err != nil {
		return "", err
	}
	return title, nil
}
