package main

import (
	"api/src/configs"
	"api/src/initializers"
	"log/slog"
	"os"
)

// TODO write the tests

var APIServerCnf configs.APIserver
var PSQLCnf configs.PSQLConfig
var SecretKey = []byte(os.Getenv("SECRET_KEY"))

func main() {
	slog.Info("Start")
	initializers.InitChiRouting()
	initializers.InitDBconnection()
	slog.Info("server is closed", "info", <-initializers.Done)
	slog.Info("End")
}