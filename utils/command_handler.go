package utils

import (
	"fmt"
	"strconv"
	"strings"

	cfg "github.com/jkeeya/toado/cfg"
	. "github.com/jkeeya/toado/interfaces"
	. "github.com/jkeeya/toado/models"
)

func CommandHandler(repository TaskRepository, user_input string) error {
	// TODO: переделать
	command, err := strconv.Atoi(user_input)
	if err != nil {
		fmt.Println(cfg.Message["expected_uint"])
	}

	switch command {
	case 0:
		fmt.Println("Эм")
	case 1:
		var task_name, task_date string
		fmt.Println("Введите задачу:")
		fmt.Scanln(&task_name)

		fmt.Println(`Введите срок дедлайна в свободном формате. Нажмите Enter чтобы пропустить`)
		fmt.Scanln(&task_date)

		repository.AddTask(&Task{
			Name:    task_name,
			ExpDate: task_date,
			Done:    false,
		})

	case 2:
		fmt.Println(cfg.Message["choose_task"], "\n", cfg.Message["promt"])
		var task_id uint
		_, err := fmt.Scanln(&task_id)
		if err != nil {
			if strings.Contains(err.Error(), "invalid syntax") {
				fmt.Println(cfg.Message["expected_uint"])
			} else {
				fmt.Println(cfg.Message["error"])
			}
		}
		repository.MarkDone(task_id)
		fmt.Println(cfg.Message["task_done"])

	case 3:
		var task_to_delete uint
		_, err := fmt.Scanln(&task_to_delete)
		if err != nil {
			if strings.Contains(err.Error(), "invalid syntax") {
				fmt.Println(cfg.Message["expected_uint"])
			} else {
				fmt.Println(cfg.Message["error"])
			}
		}

		err = repository.DeleteTask(task_to_delete)
		if err != nil {
			fmt.Println(fmt.Println(cfg.Message["error"], err))
		} else {
			fmt.Println(cfg.Message["task_deleted"])
		}
	default:
		fmt.Println(cfg.Message["wrong_option"])
	}
	return nil
}
