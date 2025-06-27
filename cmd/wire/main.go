package main

import "context"

func main() {
	userHandler := InitUserHandler()

	userHandler.RegisterUser(context.Background(), "daniel")
}
