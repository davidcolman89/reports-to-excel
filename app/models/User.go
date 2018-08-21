package models

import "time"

type User struct {
	Id 					int			`csv:"id" json:"id"`
	Nombre 				string		`csv:"nombre" json:"nombre"`
	Apellido 			string		`csv:"apellido" json:"apellido"`
	Email 				string		`csv:"email" json:"email"`
	Password 			string		`csv:"password" json:"password"`
	Dni 				int			`csv:"dni" json:"dni"`
	IdFuerza 			int			`csv:"id_fuerza" json:"id_fuerza"`
	IdPerfil 			int			`csv:"id_perfil" json:"id_perfil"`
	Username 			string		`csv:"username" json:"username"`
	RememberToken 		string		`csv:"remember_token" json:"remember_token"`
	CreatedAt 			time.Time	`csv:"created_at" json:"created_at"`
	UpdatedAt 			time.Time	`csv:"updated_at" json:"updated_at"`
	Telefono 			string		`csv:"telefono" json:"telefono"`
	ContactoLaboral 	string 		`csv:"contacto_laboral" json:"contacto_laboral"`
	Area 				string		`csv:"area" json:"area"`
	Activo 				int			`csv:"activo" json:"activo"`
}
