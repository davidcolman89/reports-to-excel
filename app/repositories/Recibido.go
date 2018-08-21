package repositories

import (
	"github.com/davidcolman89/reports-to-excel/app/models"
	"github.com/davidcolman89/reports-to-excel/app/entities"
	"os"
	"github.com/gocarina/gocsv"
	"github.com/jmoiron/sqlx"
	"github.com/davidcolman89/manager-sql-struct"
	"github.com/MakeNowJust/heredoc"
	"io"
	"encoding/csv"
)

type RecibidoRepo struct {
	Db *sqlx.DB
	Path string
}

func NewRecibidoRepo(db *sqlx.DB, path string) Recibido {
	return RecibidoRepo{ db, path}
}

func (r RecibidoRepo) Select() ([]models.ReporteRecibido, error){

	var entity []entities.ReporteRecibido
	var model []models.ReporteRecibido

	query := heredoc.Doc(`
`)
	err := r.Db.Select(&entity, query)

	if err!=nil {
		return model, err
	}

	err = sqlstruct.Marshall(entity, &model)

	if err != nil {
		return model, err
	}

	return model, nil
}

func (r RecibidoRepo) CreateCsv(users []models.ReporteRecibido) error {

	clientsFile, err := os.OpenFile(r.Path, os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		return err
	}

	defer clientsFile.Close()

	gocsv.SetCSVWriter(func(out io.Writer) *gocsv.SafeCSVWriter {
		writer := csv.NewWriter(out)
		writer.Comma = '|'
		return gocsv.NewSafeCSVWriter(writer)
	})

	return gocsv.MarshalFile(&users, clientsFile)
}
