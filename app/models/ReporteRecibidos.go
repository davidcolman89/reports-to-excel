package models

import (
	"time"
)

type ReporteRecibido struct {
	Id 						int			`csv:"id"`
	CaracterDeclaracion 	string		`csv:"caracter_declaracion"`
	Version 				string		`csv:"version"`
	TipoEstado 				string		`csv:"tipo_estado"`
	CreatedAt 				time.Time	`csv:"creado"`
	AcceptedAt 				time.Time	`csv:"aceptado"`
	ReceivedAt 				time.Time	`csv:"recibido"`
	DeletedAt 				time.Time	`csv:"eliminado"`
	Nombre 					string		`csv:"nombre"`
	Apellido 				string		`csv:"apellido"`
	Documento 				string 		`csv:"documento"`
	Fuerza 					string		`csv:"fuerza"`
}
