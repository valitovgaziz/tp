package main

import (
	"api/src/initializers"
)

func main() {
	initializers.InitChiRouting()
	initializers.InitDBconnection()
}
