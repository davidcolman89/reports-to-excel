package repositories

import "github.com/davidcolman89/reports-to-excel/app/models"

type Recibido interface {
	Select() ([]models.ReporteRecibido, error)
	CreateCsv(users []models.ReporteRecibido) error
}

type User interface {
	Select() ([]models.User, error)
	CreateCsv(users []models.User) error
}