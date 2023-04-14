package main

import (
	"github.com/mezz-ir/bookstore-clean-arch/controller"
	"github.com/mezz-ir/bookstore-clean-arch/database"
	"github.com/mezz-ir/bookstore-clean-arch/service" 

	"github.com/labstack/echo"
)

func main() {

	echoContext := echo.New()

	datalayer := database.NewBookDatalayer()
	service := service.NewBookService(datalayer))
	controller.NewBookController(echoContext, service)

	echoContext.Logger.Info(e.Start(":8070"))
}