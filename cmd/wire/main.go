package main

import (
	"errors"
	"log"
	"net/http"
)

func main() {
	app := InitApp()

	err := app.Run("localhost:9420")
	if err != nil && errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("server err: %v", err)
	}
}
