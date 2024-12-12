package taskService

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Text   string `json:"text"`
	IsDone bool   `json:"is_done"`
	Task   string `json:"task"`
}


// type Response struct {
// 	Status string `json:"status"`
// 	Message string `json:"message"`
// }
