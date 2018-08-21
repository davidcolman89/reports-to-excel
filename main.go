package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/gin-gonic/gin"
	"github.com/davidcolman89/reports-to-excel/app/repositories"
	"github.com/davidcolman89/reports-to-excel/app/services"
	"github.com/davidcolman89/reports-to-excel/config"
	"net/http"
	"fmt"
	"time"
)

func main(){
	conf := config.NewConfig("Config")
	router := gin.Default()
	router.GET("users", users)
	router.GET("recibidas", declaracionesRecibidas)
	router.Run(conf.HttpServer)
}

func users(context *gin.Context) {
	db, err := InitDb()
	HandleError(context, err)

	path := fmt.Sprintf("storage/users_%v.csv",time.Now().UnixNano())
	userRepo := repositories.NewUserRepo(db,path)
	userService := services.NewUserService(userRepo)

	users, err := userService.Select()
	HandleError(context, err)

	err = userService.CreateCsv(users)
	HandleError(context, err)

	context.File(path)

}

func declaracionesRecibidas(context *gin.Context) {
	db, err := InitDb()
	HandleError(context, err)

	fileName := fmt.Sprintf("ddjjrecibidas_%v.csv",time.Now().UnixNano())
	path := fmt.Sprint("storage/",fileName)
	repo := repositories.NewRecibidoRepo(db,path)
	service := services.NewRecibidoService(repo)

	recibidos, err := service.Select()
	HandleError(context, err)

	err = service.CreateCsv(recibidos)
	HandleError(context, err)

	context.Header("Content-type","application/octet-stream")
	contentDisposition := fmt.Sprintf("attachment; filename=%s", fileName)
	context.Header("Content-Disposition", contentDisposition)

	context.File(path)
}

func HandleError(ctx *gin.Context, err error) {
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError,err)
	}
}

func InitDb() (*sqlx.DB, error){
	conf := config.NewConfig("Config")
	dataSourceName := fmt.Sprintf("%s:%s@(%s:%s)/%s",conf.DbUser,conf.DbPass,conf.DbHost,conf.DbPort,conf.DbBbdd)
	return sqlx.Connect("mysql", dataSourceName)
}