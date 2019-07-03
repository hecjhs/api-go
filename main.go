package main

import (
	handlers "github.com/hecjhs/api-go/api/handlers"
	server "github.com/hecjhs/api-go/api/server"
	utils "github.com/hecjhs/api-go/api/utils"
)

func main() {
	utils.LoadEnv()
	app := server.SetUp()
	handlers.HandlerRedirection(app)
	server.RunServer(app)
}
