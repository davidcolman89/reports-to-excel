# reports sql queries to excel(CSV) with enconding Windows-1252
## How to use?
 * At the root of the project:
 
  1- create 'config.json'
  ```json
  {
    "HttpServer":":8080",
    "dbHost":"localhost",
    "dbPort":"9999",
    "dbUser":"root",
    "dbPass":"123456",
    "dbBbdd":"name_of_databases",
    "csvDelimiter":";"
  }
  ```
  
  2- Create folder 'queries' that contains the *.sql files
  
  3- Set your 'entities' files on 'app/entities/...', that is your input data. Example:
```go
package entities

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

type ExampleInputData struct {
	Id 						sql.NullInt64		`db:"id"`
	Version 				sql.NullString		`db:"version"`
	CreatedAt 				mysql.NullTime		`db:"created_at"`
	AcceptedAt 				mysql.NullTime		`db:"accepted_at"`
	ReceivedAt 				mysql.NullTime		`db:"received_at"`
	DeletedAt 				mysql.NullTime		`db:"deleted_at"`
	Name 					sql.NullString		`db:"name"`
	Fullname 				sql.NullString		`db:"fullname"`
  ...
}
```

  4- Set your 'models' files on 'app/models/...', that is your output/csv data Example:
```go
package models

import (
	"time"
)

type ExampleOutputData struct {
	Id 						int			`csv:"id"`
	Version 				string		`csv:"version"`
	CreatedAt 				time.Time	`csv:"creado"`
	AcceptedAt 				time.Time	`csv:"aceptado"`
	ReceivedAt 				time.Time	`csv:"recibido"`
	DeletedAt 				time.Time	`csv:"eliminado"`
	Name 				string		`csv:"name"`
  Fullname 				string		`csv:"apellido"`
	...
}
```

  5- the run 'main.go' 
  ```sh
  go run main.go
  ```
