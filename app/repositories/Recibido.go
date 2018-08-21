package repositories

import (
	"github.com/davidcolman89/reports-to-excel/app/models"
	"github.com/davidcolman89/reports-to-excel/app/entities"
	"os"
	"github.com/gocarina/gocsv"
	"github.com/jmoiron/sqlx"
	"github.com/davidcolman89/manager-sql-struct"
	"io"
	"encoding/csv"
	"io/ioutil"
	"golang.org/x/text/encoding/charmap"
	"fmt"
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

	err := r.getDataByFuerza(&entity, "pfa")

	if err!=nil {
		return model, err
	}

	err = r.getDataByFuerza(&entity, "gna")

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
		//Create new writer that accept encoding 'Windows-1252'
		writerWindows1552 := charmap.Windows1252.NewEncoder().Writer(out)

		writer := csv.NewWriter(writerWindows1552)
		writer.Comma = ';'

		return gocsv.NewSafeCSVWriter(writer)
	})

	return gocsv.MarshalFile(&users, clientsFile)
}

func (r RecibidoRepo) getDataByFuerza(entity *[]entities.ReporteRecibido, fuerza string) error {

	filename := fmt.Sprintf("./queries/recibidos/%v.sql",fuerza)
	sqlFile, err := ioutil.ReadFile(filename)

	if err!=nil {
		return err
	}

	query := string(sqlFile)

	return r.Db.Select(entity, query)
}