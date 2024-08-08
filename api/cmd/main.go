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
	initializers.InitDBconnection()
	initializers.InitChiRouting()
	slog.Info("End")
}