package main

import (
	"fmt"
	"github.com/claravelita/majoo-test/api"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	fmt.Println("application has started...")
	newEcho := echo.New()
	serverAPI := api.NewServer(newEcho)

	serverAPI.InitializeServer()
}
