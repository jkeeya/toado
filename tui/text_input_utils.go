package tui

import (
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	. "github.com/jkeeya/toado/models"
)

// Смена полей ввода названия и даты таска по нажатию Tab
func serveChangeTextInput(m *model, msg tea.KeyMsg) {
	if !m.taskToDeleteInput.Focused() && !m.markDoneInput.Focused() {
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
}

func serveAddTask(m *model) {
	m.isAwaitingInput = false
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

func serveDeleteTask(m *model) error {
	taskID := m.taskToDeleteInput.Value()
	id, err := strconv.ParseUint(strings.TrimSpace(taskID), 10, 64)
	if err != nil {
		return err
	}
	m.app.Repository.DeleteTask(uint(id))
	m.updateTaskList()
	m.isAwaitingInput = false
	m.taskToDeleteInput.Reset()
	m.taskToDeleteInput.Blur()
	return nil
}

func serveMarkDone(m *model) error {
	taskID := m.markDoneInput.Value()
	id, err := strconv.ParseUint(strings.TrimSpace(taskID), 10, 64)
	if err != nil {
		return err
	}
	m.app.Repository.MarkDone(uint(id))
	m.updateTaskList()
	m.isAwaitingInput = false
	m.markDoneInput.Reset()
	m.markDoneInput.Blur()
	return nil
}
