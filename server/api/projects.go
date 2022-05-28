package api

type Project struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Color       string `json:"color"`
	Importance  string `json:"importance"`
	ProjectId   string `json:"projectId"`
}
