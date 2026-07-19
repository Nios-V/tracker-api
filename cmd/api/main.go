package main

import (
	"github.com/Nios-V/tracker-api/internal/database"
	"github.com/Nios-V/tracker-api/internal/handlers"
	"github.com/Nios-V/tracker-api/internal/repository"
	"github.com/Nios-V/tracker-api/internal/router"
	"github.com/Nios-V/tracker-api/internal/services"
)

func main() {
	database.Connect()

	repos := repository.NewRepositories(database.DB)
	svcs := services.NewServices(repos)
	hdlrs := handlers.NewHandlers(svcs)

	r := router.SetupRouter(hdlrs)
	r.Run(":8080")
}
