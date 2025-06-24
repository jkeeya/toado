package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"toado/db"
	. "toado/interfaces"
)

func main() {
	// Инициализируем БД прямо в main, позже создадим интерфейс для хранилищ
	DB := db.InitDB()
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
		// Считываем ввод пользователя
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка, попробуйте ещё раз")
			continue
		}
		// Подчищаем ввод от символов переноса и пробелов
		input = input[:len(input)-1]
		// TODO: нужно?
		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Сосал?")
			continue
		}

		switch number {
		case 0:
			fmt.Println("Выход из программы... Спасибо что вы с нами!")
			return
		case 1:
			fmt.Println("Введите задачу: ")
			//
			// TODO: работать с датами а не строками
			fmt.Println(`
			Введите срок дедлайна в свободном формате.
			Нажмите Enter чтобы пропустить`)
			//
			// repository.AddTask()
		case 2:
			// fmt.Println("Выполняется действие для клавиши 2")
		case 3:
			// fmt.Println("Выполняется действие для клавиши 3")
		default:
			fmt.Println("Эта опция ещё в разработке")
		}

	}
}
