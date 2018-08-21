package services

import (
	"github.com/davidcolman89/reports-to-excel/app/models"
	"github.com/davidcolman89/reports-to-excel/app/repositories"
)

type UserService struct {
	Repo repositories.User
}


func NewUserService(repo repositories.User) User{
	return UserService{repo}
}

func (s UserService) CreateCsv(users []models.User) error {
	return s.Repo.CreateCsv(users)
}

func (s UserService) Select()  ([]models.User, error){
	return s.Repo.Select()
}
