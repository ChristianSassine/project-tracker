package api

import (
	"time"
)

type Project struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type Task struct {
	Id           int       `json:"id"`
	Title        string    `json:"title"`
	State        string    `json:"state"`
	Description  string    `json:"description"`
	CreationDate time.Time `json:"creationDate"`
	ProjectId    string    `json:"projectId"`
}

type TaskPatchRequest struct {
	NewState      string `json:"newState"`
	PreviousIndex int    `json:"previousIndex"`
	CurrentIndex  int    `json:"currentIndex"`
	TaskId        int    `json:"taskId"`
}

type TaskStats struct {
	TodoTasks    int64 `json:"todoTasks"`
	OngoingTasks int64 `json:"ongoingTasks"`
	DoneTasks    int64 `json:"doneTasks"`
}

type Log struct {
	Logger string    `json:"logger"`
	Type   string    `json:"type"`
	Args   []string  `json:"arguments"`
	Date   time.Time `json:"date"`
}
