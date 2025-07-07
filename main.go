package main

import (
	"github.com/jkeeya/toado/cfg"
	"github.com/jkeeya/toado/db"
	tui "github.com/jkeeya/toado/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Инициализируем БД прямо в main, позже создадим интерфейс для хранилищ (если понадобится)
	var app cfg.App
	app.DB = db.InitDB("data.db")
	app.Repository = &db.SQLiteTaskRepository{DB: app.DB}

	tui_model := tui.NewTeaModel(app)
	p := tea.NewProgram(tui_model)
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
