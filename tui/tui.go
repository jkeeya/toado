package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/jkeeya/toado/cfg"
	. "github.com/jkeeya/toado/models"
	"github.com/jkeeya/toado/utils"
)

type model struct {
	app        cfg.App
	taskList   []Task // Текущий список задач
	options    string // Список доступных действий
	lastResult string // Результат последнего действа
	command    string // Вводимая юзером команда
}

func NewTeaModel(app cfg.App) *model {
	taskList := app.Repository.GetTasks()
	options := cfg.Message["menu"]
	return &model{
		app:        app,
		taskList:   taskList,
		options:    options,
		lastResult: "",
		command:    "",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyRunes:
			if len(msg.String()) == 1 {
				m.command += msg.String()
			}
		case tea.KeyCtrlC:
			fmt.Println(cfg.Message["exit"])
			return m, tea.Quit
		case tea.KeyEnter:
			// TODO: ?
			utils.CommandHandler(m.app.Repository, m.command)
			m.command = ""
		}
	}
	return m, nil
}

func (m model) View() string {
	// Постоянная часть: список задач
	// TODO: когда список задач большой, отображаться должны не все, и д. б. возможность листать список.
	// Придумать кому поручить фильтрацию, туе или утилитам
	taskList := "Список задач:\n" + utils.TasksToString(m.app.Repository.GetTasks())
	menu := cfg.Message["menu"]

	// Динамическая часть: результат и ввод
	inputLine := fmt.Sprintf("\n> %s", m.command)
	resultLine := fmt.Sprintf("\n%s\n", m.lastResult)

	return taskList + menu + resultLine + inputLine
}
