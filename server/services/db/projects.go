package db

import (
	"BugTracker/api"
	"strings"
)

func (db *DB) GetUserProjects(username string) (*[]api.Project, error) {

	projects := []api.Project{}
	project := &api.Project{}

	username = strings.ToLower(username)
	rows, err := db.DB.Query(
		`SELECT projects.id, projects.title FROM projects 
	INNER JOIN users_projects ON projects.id = users_projects.project_id
	INNER JOIN users ON users.id = users_projects.user_id
	WHERE users.username = $1`,
		username)

	if err != nil {
		return &[]api.Project{}, err
	}

	for rows.Next() {
		if err := rows.Scan(&project.Id, &project.Title); err != nil {
			return &[]api.Project{}, err
		}
		projects = append(projects, *project)
	}

	return &projects, nil
}

func (db *DB) CreateProject(userId int, title string) error {
	var projectId int

	row := db.DB.QueryRow("INSERT INTO projects (title) VALUES ($1) RETURNING id", title)

	if err := row.Scan(&projectId); err != nil {
		return err
	}

	if _, err := db.DB.Exec("INSERT INTO users_projects VALUES ($1, $2)", userId, projectId); err != nil {
		return err
	}

	return nil
}

func (db *DB) GetTasks(userId int, projectId int) (*[]api.Task, error) {
	tasks := []api.Task{}
	task := api.Task{}

	rows, err := db.DB.Query(`
	SELECT tasks.id, tasks.title, tasks.description, tasks.type, tasks.color, tasks.importance FROM tasks 
	INNER JOIN projects ON tasks.project_id = projects.id
	INNER JOIN users_projects ON projects.id = users_projects.project_id
	INNER JOIN users ON users.id = users_projects.user_id
	WHERE users.id = $1 AND projects.id = $2`, userId, projectId)

	if err != nil {
		return &[]api.Task{}, err
	}

	for rows.Next() {
		if err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.Type, &task.Color, &task.Importance); err != nil {
			return &[]api.Task{}, err
		}
		tasks = append(tasks, task)
	}

	return &tasks, nil
}
