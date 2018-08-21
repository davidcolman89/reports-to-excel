package repositories

import (
	"github.com/davidcolman89/reports-to-excel/app/models"
	"os"
	"github.com/gocarina/gocsv"
	"github.com/davidcolman89/reports-to-excel/app/entities"
	"github.com/jmoiron/sqlx"
	"github.com/davidcolman89/manager-sql-struct"
)

type UserRepo struct {
	Path string
	Db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB, path string) User {
	return UserRepo{path, db}
}

func (r UserRepo) CreateCsv(users []models.User) error {

	clientsFile, err := os.OpenFile(r.Path, os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		return err
	}

	defer clientsFile.Close()

	return gocsv.MarshalFile(&users, clientsFile)
}


func (u UserRepo) Select() ([]models.User, error){

	var entity []entities.User
	var model []models.User

	err := u.Db.Select(&entity, "")

	if err!=nil {
		return model, err
	}

	err = sqlstruct.Marshall(entity, &model)

	if err != nil {
		return model, err
	}

	return model, nil
}