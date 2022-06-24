package api

import (
	"time"
)

type Project struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Password string `json:"password"`
}

type ProjectJoinRequest struct {
	Id       int    `json:"id"`
	Password string `json:"password"`
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

type Comment struct {
	Commenter string    `json:"commenter"`
	Content   string    `json:"content"`
	Date      time.Time `json:"date"`
}

type CommentAddRequest struct {
	Content string `json:"content"`
}
