package cfg

var Message = map[string]string{
	"hello":         "Hello!",
	"exit":          "Выход из программы... Спасибо что вы с нами!!",
	"promt":         "> ",
	"completed":     "✔",
	"not_completed": "✘",

	// Ошибки
	"error":         "Ошибка...",
	"expected_uint": "Введите целое число",
	"wrong_option":  "Эта опция ещё в разработке",

	"menu": (`
1 - Добавить новую задачу
2 - Отметить задачу выполненной
3 - Удалить задачу
	`),

	// Команды
	"choose_task": "Выберите задачу",

	// Результаты выполнения
	"task_done":    "Так держать!",
	"task_deleted": "Задача удалена",
}
