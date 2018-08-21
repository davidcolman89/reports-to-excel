package services

import (
	"github.com/davidcolman89/reports-to-excel/app/models"
	"github.com/davidcolman89/reports-to-excel/app/repositories"
)

type RecibidoService struct {
	Repo repositories.Recibido
}

func NewRecibidoService(repo repositories.Recibido) Recibido{
	return RecibidoService{repo}
}

func (s RecibidoService) CreateCsv(recibidos []models.ReporteRecibido) error {
	return s.Repo.CreateCsv(recibidos)
}

func (s RecibidoService) Select()  ([]models.ReporteRecibido, error){
	return s.Repo.Select()
}
