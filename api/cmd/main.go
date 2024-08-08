package main

import (
	"api/src/configs"
	"api/src/initializers"
	"log/slog"
)

var APIServerCnf configs.APIserver
var PSQLCnf configs.PSQLConfig

func main() {
	slog.Info("Start")
	initializers.InitChiRouting()
	initializers.InitDBconnection()
	slog.Info("server is closed", "info", <-initializers.Done)
	slog.Info("End")
}