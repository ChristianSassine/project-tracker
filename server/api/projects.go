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

type TaskPositionRequest struct {
	PreviousIndex int `json:"previousIndex"`
	CurrentIndex  int `json:"currentIndex"`
	TaskId        int `json:"taskId"`
}
