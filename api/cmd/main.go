package main

import (
	"api/src/initializers"
	"log/slog"
)

func main() {
	slog.Info("Start")
	initializers.InitChiRouting()
	initializers.InitDBconnection()
	slog.Info("End")
}