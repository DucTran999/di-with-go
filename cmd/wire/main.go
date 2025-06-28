package main

func main() {
	app := InitApp()

	app.router.Run("localhost:9420")
}
