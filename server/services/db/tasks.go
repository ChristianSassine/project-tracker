package db

import "BugTracker/api"

func (db *DB) GetTasks(userId int, projectId int) (*[]api.Task, error) {
	tasks := []api.Task{}
	task := api.Task{}

	rows, err := db.DB.Query(`
	SELECT tasks.id, tasks.title, tasks.description, tasks.type FROM tasks 
	INNER JOIN projects ON tasks.project_id = projects.id
	INNER JOIN users_projects ON projects.id = users_projects.project_id
	INNER JOIN users ON users.id = users_projects.user_id
	WHERE users.id = $1 AND projects.id = $2`, userId, projectId)

	if err != nil {
		return &[]api.Task{}, err
	}

	for rows.Next() {
		if err := rows.Scan(&task.Id, &task.Title, &task.Description); err != nil {
			return &[]api.Task{}, err
		}
		tasks = append(tasks, task)
	}

	return &tasks, nil
}

func (db *DB) AddTask(task *api.Task, projectId int) error {

	_, err := db.DB.Exec(`
	INSERT INTO tasks (title, state, description, creation_date, project_id)
	VALUES ($1, $2, $3, $4, $5)`, task.Title, task.State, task.Description, task.CreationDate, projectId)
	if err != nil {
		return err
	}

	return nil
}
