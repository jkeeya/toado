package tui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/jkeeya/toado/cfg"
	. "github.com/jkeeya/toado/models"
	"github.com/jkeeya/toado/utils"
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

	taskNameInput     textinput.Model
	deadlineInput     textinput.Model
	taskToDeleteInput textinput.Model

	currentInput    int  // Индекс текущего активного поля (0 - название, 1 - дедлайн)
	isAwaitingInput bool // Флаг ожидания дополнительного ввода

}

func (m *model) updateTaskList() {
	m.taskList = m.app.Repository.GetTasks()
}

func NewTeaModel(app cfg.App) *model {
	taskNameInput := textinput.New()
	deadlineInput := textinput.New()
	taskToDeleteInput := textinput.New()

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
		app:               app,
		taskList:          taskList,
		options:           options,
		taskNameInput:     taskNameInput,
		deadlineInput:     deadlineInput,
		taskToDeleteInput: taskToDeleteInput,
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

		// Режим ввода: работа с текстовыми полями
		if m.isAwaitingInput {
			switch msg.Type {
			case tea.KeyTab, tea.KeyShiftTab:
				if !m.taskToDeleteInput.Focused() {
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
				}
			// Добавление задачи
			case tea.KeyEnter:
				if m.currentInput == 1 { // Последнее поле
					m.isAwaitingInput = false
					// TODO: мб надо изменить уровень абстракции
					m.app.Repository.AddTask(&Task{
						Name:    m.taskNameInput.Value(),
						ExpDate: m.deadlineInput.Value(),
					})
					m.updateTaskList()
					m.taskNameInput.Reset()
					m.deadlineInput.Reset()
					m.taskNameInput.Blur()
					m.deadlineInput.Blur()
				}
				// Удаление задачи
				if m.taskToDeleteInput.Focused() {
					taskID := m.taskToDeleteInput.Value()
					id, err := strconv.ParseUint(strings.TrimSpace(taskID), 10, 64)
					if err != nil {
						return m, nil
					}
					m.app.Repository.DeleteTask(uint(id))
					m.updateTaskList()
					m.isAwaitingInput = false
					m.taskToDeleteInput.Reset()
					m.taskToDeleteInput.Blur()
				}

			}

			// Работа с меню
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
	case requestTaskDeleteMsg:
		m.isAwaitingInput = true
		m.taskToDeleteInput.Focus()
		return m, nil
	}

	m.options, cmd = m.options.Update(msg)
	m.taskNameInput, cmd = m.taskNameInput.Update(msg)
	m.deadlineInput, cmd = m.deadlineInput.Update(msg)
	m.taskToDeleteInput, cmd = m.taskToDeleteInput.Update(msg)

	return m, cmd
}

func (m model) View() string {
	if m.isAwaitingInput {
		if m.taskToDeleteInput.Focused() {
			view := "Введите ID задачи для удаления:" + m.taskToDeleteInput.View() + "\n\n"
			return utils.TasksToString(m.taskList) + view
		} else {
			view := "Название задачи:\n" + m.taskNameInput.View() + "\n\n"
			view += "Дедлайн:\n" + m.deadlineInput.View() + "\n\n"
			view += "Используйте Tab для переключения между полями. Нажмите Enter для завершения."
			return view
		}
	}

	// Отображение списка
	return utils.TasksToString(m.taskList) + m.options.View()
}
