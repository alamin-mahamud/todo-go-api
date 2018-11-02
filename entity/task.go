package entity

type Task struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Done bool `json:"done"`
}

type TaskCollection struct {
	Tasks []Task `json:"items"`
}