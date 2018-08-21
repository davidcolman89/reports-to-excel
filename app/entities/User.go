package entities

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

type User struct {
	Id 					sql.NullInt64		`db:"id"`
	Nombre 				sql.NullString		`db:"nombre"`
	Apellido 			sql.NullString		`db:"apellido"`
	Email 				sql.NullString		`db:"email"`
	Password 			sql.NullString		`db:"password"`
	Dni 				sql.NullInt64		`db:"dni"`
	IdFuerza 			sql.NullInt64		`db:"id_fuerza"`
	IdPerfil 			sql.NullInt64		`db:"id_perfil"`
	Username 			sql.NullString		`db:"username"`
	RememberToken 		sql.NullString		`db:"remember_token"`
	CreatedAt 			mysql.NullTime		`db:"created_at"`
	UpdatedAt 			mysql.NullTime		`db:"updated_at"`
	Telefono 			sql.NullString		`db:"telefono"`
	ContactoLaboral 	sql.NullString 		`db:"contacto_laboral"`
	Area 				sql.NullString		`db:"area"`
	Activo 				sql.NullInt64		`db:"activo"`
}

