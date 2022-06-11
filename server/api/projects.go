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

type Log struct {
	Logger  string    `json:"initiator"`
	Content string    `json:"content"`
	Date    time.Time `json:"date"`
}
