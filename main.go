package main

import (
	"fmt"

	"github.com/JerryJeager/will-be-there-backend/cmd"
	"github.com/JerryJeager/will-be-there-backend/config"
)


func main() {
	config.LoadEnv()
	config.ConnectToDB()
	fmt.Println("env and database initializaed successfully...")
	fmt.Println("starting to the will be there server...")
	cmd.ExecuteApiRoutes()
}
