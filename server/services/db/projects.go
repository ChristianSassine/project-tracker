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

func (db *DB) CreateProject(userId int, title string, password string) (*api.Project, error) {
	var project api.Project

	if err := db.DB.QueryRow("INSERT INTO projects (title, password) VALUES ($1, $2) RETURNING id, title", title, password).Scan(&project.Id, &project.Title); err != nil {
		return &api.Project{}, err
	}

	if _, err := db.DB.Exec("INSERT INTO users_projects VALUES ($1, $2)", userId, project.Id); err != nil {
		return &api.Project{}, err
	}

	return &project, nil
}

func (db *DB) DeleteProject(projectId int) error {
	if _, err := db.DB.Exec(`DELETE FROM projects WHERE id = $1`, projectId); err != nil {
		return err
	}

	return nil
}

func (db *DB) GetProjectPassword(projectId int) (string, error) {
	var password string
	if err := db.DB.QueryRow(`SELECT password FROM projects WHERE id = $1`, projectId).Scan(&password); err != nil {
		return password, err
	}

	return password, nil
}

func (db *DB) GetProjectTitle(projectId int) (string, error) {
	var title string
	if err := db.DB.QueryRow(`SELECT title FROM projects WHERE id = $1`, projectId).Scan(&title); err != nil {
		return title, err
	}

	return title, nil
}

func (db *DB) AddUserToProject(userId int, projectId int) error {
	if _, err := db.DB.Exec(`INSERT INTO users_projects (user_id, project_id) VALUES ($1, $2)`, userId, projectId); err != nil {
		return err
	}

	return nil
}

func (db *DB) IsUserInProject(userId int, projectId int) bool {

	var userIdCheck int
	if err := db.DB.QueryRow(`SELECT user_id FROM users_projects WHERE user_id = $1 AND project_id = $2`, userId, projectId).Scan(&userIdCheck); err != nil {
		return false
	}

	return true
}
