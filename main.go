package main

import (
	"os"

	"jwt-project/common/env"
	"jwt-project/controller"
	"jwt-project/repository"
	"jwt-project/service"

	"jwt-project/routes"
)

func main() {
	env.Load()
	router := setupAllDependencies()

	port := os.Getenv("PORT")
	url := ":" + port

	router.Run(url)

}

func setupAllDependencies() routes.Router {
	repo := repository.NewRepository()
	serv := service.NewService(repo)
	control := controller.NewController(serv)

	router := routes.NewRouter(control)
	router.Setup()

	return router
}
