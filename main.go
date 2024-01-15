package main

import server "github.com/Rockuse/go-ecommerce/src"

func main() {
	app := server.NewServer()
	app.Run()
}
