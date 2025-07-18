package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/jkeeya/toado/cfg"
	. "github.com/jkeeya/toado/models"
	// "github.com/jkeeya/toado/utils"
)

type listItem struct {
	title       string
	description string
	action      func() tea.Cmd
}

func (i listItem) Title() string       { return i.title }
func (i listItem) Description() string { return i.description }
func (i listItem) FilterValue() string { return i.title }

type model struct {
	app      cfg.App
	taskList []Task     // Текущий список задач
	options  list.Model // Список доступных действий

	taskNameInput textinput.Model
	deadlineInput textinput.Model
	currentInput  int // Индекс текущего активного поля (0 - название, 1 - дедлайн)
	//lastResult string // Результат последнего действа
	//UserInput  string // Вводимая юзером команда
	isAwaitingInput bool // Флаг ожидания дополнительного ввода

}

func NewTeaModel(app cfg.App) *model {
	ti1 := textinput.New()
	ti2 := textinput.New()

	// TODO: Вынести инициализацию меню отдельно
	items := []list.Item{
		listItem{
			title:       "Добавить задачу",
			description: "Добавить задачу",
			action: func() tea.Cmd {
				return requestTaskInput()
			},
		},
		listItem{
			title:       "Удалить задачу",
			description: "Удалить задачу",
			action: func() tea.Cmd {
				return requestTaskDelete()
			},
		},
	}

	options := list.New(items, list.NewDefaultDelegate(), 20, 10)
	options.Title = "Меню"

	taskList := app.Repository.GetTasks()
	return &model{
		app:           app,
		taskList:      taskList,
		options:       options,
		taskNameInput: ti1,
		deadlineInput: ti2,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:

		if msg.Type == tea.KeyCtrlC {
			fmt.Println(cfg.Message["exit"])
			return m, tea.Quit
		}

		if m.isAwaitingInput {
			switch msg.Type {
			case tea.KeyTab, tea.KeyShiftTab:
				// Переключение между полями
				if msg.Type == tea.KeyTab {
					m.currentInput = (m.currentInput + 1) % 2
				} else if msg.Type == tea.KeyShiftTab {
					m.currentInput = (m.currentInput - 1 + 2) % 2
				}

				m.taskNameInput.Blur()
				m.deadlineInput.Blur()
				if m.currentInput == 0 {
					m.taskNameInput.Focus()
				} else {
					m.deadlineInput.Focus()
				}
			case tea.KeyEnter:
				if m.currentInput == 1 { // Последнее поле
					m.isAwaitingInput = false
					// TODO: мб надо изменить уровень абстракции
					m.app.Repository.AddTask(&Task{
						Name:    m.taskNameInput.Value(),
						ExpDate: m.deadlineInput.Value(),
					})
					m.taskNameInput.Reset()
					m.deadlineInput.Reset()
					// TODO: ??????
					m.taskNameInput.Blur()
					m.deadlineInput.Blur()
				}
			}
		} else {
			switch msg.Type {
			case tea.KeyEnter:
				selectedItem := m.options.SelectedItem().(listItem)
				cmd = selectedItem.action()
				return m, cmd

			}

		}
	case requestTaskInputMsg:
		m.isAwaitingInput = true
		m.taskNameInput.Focus()
		return m, nil
	}
	if !m.isAwaitingInput {
		m.options, cmd = m.options.Update(msg)
	} else { // TODO: что за пиздец можно оптимальнее?
		if m.currentInput == 0 {
			m.taskNameInput, cmd = m.taskNameInput.Update(msg)
		} else {
			m.deadlineInput, cmd = m.deadlineInput.Update(msg)
		}
	}

	return m, cmd
}

func (m model) View() string {
	if m.isAwaitingInput {
		// Отображение текстовых полей
		view := "Введите данные для задачи:\n\n"
		view += "Название задачи:\n" + m.taskNameInput.View() + "\n\n"
		view += "Дедлайн:\n" + m.deadlineInput.View() + "\n\n"
		view += "Используйте Tab для переключения между полями. Нажмите Enter для завершения."
		return view
	}

	// Отображение списка
	return m.options.View()
}
