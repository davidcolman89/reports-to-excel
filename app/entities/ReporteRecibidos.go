package entities

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

type ReporteRecibido struct {
	Id 						sql.NullInt64		`db:"id"`
	CaracterDeclaracion 	sql.NullString		`db:"descripcion"`
	Version 				sql.NullString		`db:"version"`
	TipoEstado 				sql.NullString		`db:"tipo_estado"`
	CreatedAt 				mysql.NullTime		`db:"created_at"`
	AcceptedAt 				mysql.NullTime		`db:"accepted_at"`
	ReceivedAt 				mysql.NullTime		`db:"received_at"`
	DeletedAt 				mysql.NullTime		`db:"deleted_at"`
	Nombre 					sql.NullString		`db:"nombres"`
	Apellido 				sql.NullString		`db:"apellido"`
	Documento 				sql.NullString 		`db:"documento"`
	Fuerza 					sql.NullString		`db:"fuerza"`
}

