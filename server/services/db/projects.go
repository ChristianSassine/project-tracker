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
