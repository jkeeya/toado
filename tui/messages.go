package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Запрос ввода таска
type requestTaskInputMsg struct{}

func requestTaskInput() tea.Cmd {
	return func() tea.Msg {
		return requestTaskInputMsg{}
	}
}

// Удаление таска
type requestTaskDeleteMsg struct{}

func requestTaskDelete() tea.Cmd {
	return func() tea.Msg {
		return requestTaskDeleteMsg{}
	}
}
