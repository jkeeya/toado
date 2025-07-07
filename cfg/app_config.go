package cfg

import (
	. "github.com/jkeeya/toado/interfaces"

	gorm "gorm.io/gorm"
)

type App struct {
	Repository TaskRepository
	DB         *gorm.DB
	TeaModel   struct{}
}
