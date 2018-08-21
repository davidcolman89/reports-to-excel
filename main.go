package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/gin-gonic/gin"
	"github.com/davidcolman89/reports-to-excel/app/repositories"
	"github.com/davidcolman89/reports-to-excel/app/services"
	"net/http"
	"fmt"
	"time"
)

func main(){

	router := gin.Default()
	router.GET("users", users)
	router.GET("recibidos", reporteRecibidos)
	router.Run()

}


func users(context *gin.Context) {

	db, err := sqlx.Connect("mysql", "root:123456@(localhost:9999)/sdj_20180630")
	HandleError(context, err)
	userRepo := repositories.NewUserRepo(db,"storage/users.csv")
	userService := services.NewUserService(userRepo)

	users, err := userService.Select()
	HandleError(context, err)

	err = userService.CreateCsv(users)
	HandleError(context, err)

}

func reporteRecibidos(context *gin.Context) {
	db, err := sqlx.Connect("mysql", "root:123456@(localhost:9999)/sdj_20180630")
	HandleError(context, err)

	path := fmt.Sprint("storage/recibidos_", time.Now().UnixNano(),".csv")
	repo := repositories.NewRecibidoRepo(db,path)
	service := services.NewRecibidoService(repo)
	recibidos, err := service.Select()

	HandleError(context, err)

	err = service.CreateCsv(recibidos)

	HandleError(context, err)
}

func HandleError(ctx *gin.Context, err error) {
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError,err)
	}
}