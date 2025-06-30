package main

import (
	"fmt"

	"github.com/jkeeya/toado/db"
	. "github.com/jkeeya/toado/interfaces"
	"github.com/jkeeya/toado/models"
)

func main() {
	// Инициализируем БД прямо в main, позже создадим интерфейс для хранилищ
	DB := db.InitDB("data.db")
	var repository TaskRepository
	repository = &db.SQLiteTaskRepository{DB: DB}

	for {
		fmt.Println(`
    1 - Добавить новую задачу
    2 - Отметить задачу выполненной
    3 - Удалить задачу
    4 - Показать все задачи
    0 - Выход
    `)
		fmt.Print("Выберите опцию: ")

		var option int
		_, err := fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Ошибка ввода. Попробуйте снова.")
			continue
		}

		switch option {
		case 0:
			fmt.Println("Выход из программы... Спасибо что вы с нами!")
			return
		case 1:
			var task_name, task_date string
			fmt.Println("Введите задачу:")
			fmt.Scanln(&task_name)

			fmt.Println(`Введите срок дедлайна в свободном формате. Нажмите Enter чтобы пропустить`)
			fmt.Scanln(&task_date)

			repository.AddTask(&models.Task{
				Name:    task_name,
				ExpDate: task_date,
				Done:    false,
			})
		case 2:
			fmt.Println("Выберите задачу:")
			var user_input uint
			_, err := fmt.Scanln(&user_input)
			if err != nil {
				fmt.Println("Ошибка ввода. Попробуйте снова.")
				continue
			}
			repository.MarkDone(user_input)
			fmt.Println("Так держать!")
		case 3:
			fmt.Println("Выберите задачу для удаления:")
			tasks, _ := repository.GetTasks()
			for _, task := range tasks {
				fmt.Printf("%d: %s\n", task.ID, task.Name)
			}

			var task_to_delete uint
			_, err := fmt.Scanln(&task_to_delete)
			if err != nil {
				fmt.Println("Ошибка ввода. Попробуйте снова.")
				continue
			}

			err = repository.DeleteTask(task_to_delete)
			if err != nil {
				fmt.Println("Ошибка при удалении задачи:", err)
			} else {
				fmt.Println("Задача удалена.")
			}
		case 4:
			tasks, _ := repository.GetTasks()
			for _, task := range tasks {
				status := "не выполнено"
				if task.Done {
					status = "выполнено"
				}
				// TODO: не нрав.. А если изменятся поля? Надо стандартизировать
				fmt.Printf("ID: %d, Задача: %s, Срок: %s, Статус: %s\n",
					task.ID, task.Name, task.ExpDate, status)
			}
		default:
			fmt.Println("Эта опция ещё в разработке")
		}
	}
}
