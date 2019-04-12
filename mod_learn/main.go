package main

import (
	"mylib"
)

func main() {
	// Set the router as the default one shipped with Gin
	router := mylib.Serve()
	router.Run(":3000")
}