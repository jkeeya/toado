package utils

import (
	"fmt"

	"github.com/jkeeya/toado/cfg"
	. "github.com/jkeeya/toado/models"
)

/*
Перевод слайса задач в строковую форму.
*/
func TasksToString(tasks []Task) string {
	var result string
	for _, task := range tasks {
		status := cfg.Message["not_completed"]
		if task.Done {
			status = cfg.Message["completed"]
		}
		result += fmt.Sprintf("%d) %s, %s, %s\n",
			task.ID, task.Name, task.ExpDate, status)
	}
	return result
}
