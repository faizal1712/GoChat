package main

import (
	"contact/routes"
)

func main() {
	routes.CreateRouter()
	routes.InititalizeRoutes()
	routes.StartServer()
}
