package main

import (
	server "github.com/Rockuse/go-ecommerce/src"
	"github.com/Rockuse/go-ecommerce/src/app/core/config"
)

func main() {
	db := config.DBInit()
	app := server.NewServer(db)
	app.Run()
}
