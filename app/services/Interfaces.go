package services

import "github.com/davidcolman89/reports-to-excel/app/models"

type User interface {
	Select() ([]models.User, error)
	CreateCsv(users []models.User) error
}

type Recibido interface {
	Select() ([]models.ReporteRecibido, error)
	CreateCsv(users []models.ReporteRecibido) error
}
