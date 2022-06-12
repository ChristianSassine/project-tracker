package db

import (
	"BugTracker/api"
	log "BugTracker/utilities"
	"time"

	"github.com/lib/pq"
)

// Log with date specified
func (db *DB) AddLogWithDate(loggerId int, projectId int, date time.Time, logType string, logArgs ...string) error {
	_, err := db.DB.Exec(`
	INSERT INTO histories (user_id, project_id, date, type, arguments) 
	VALUES ($1, $2, $3, $4, $5)`, loggerId, projectId, date, logType, pq.Array(logArgs))

	return err
}

// Log with date unspecified (uses the current date)
func (db *DB) AddLog(loggerId int, projectId int, logType string, logArgs ...string) {
	if err := db.AddLogWithDate(loggerId, projectId, time.Now().Local(), logType, logArgs...); err != nil {
		log.PrintError("Failed to add the log of type '", logType, "'. For the following error: ", err)
	}
}

func (db *DB) GetAllLogs(projectId int) (*[]api.Log, error) {
	logs := []api.Log{}
	log := api.Log{}

	rows, err := db.DB.Query(`
	SELECT histories.date, histories.type, histories.arguments, users.username FROM histories 
	INNER JOIN users ON histories.user_id = users.id
	WHERE project_id = $1 ORDER BY histories.date DESC`, projectId)

	if err != nil {
		return &[]api.Log{}, err
	}

	for rows.Next() {
		if err := rows.Scan(&log.Date, &log.Type, pq.Array(&log.Args), &log.Logger); err != nil {
			return &[]api.Log{}, err
		}
		logs = append(logs, log)
	}
	return &logs, err
}
