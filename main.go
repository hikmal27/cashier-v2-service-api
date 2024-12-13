package main

import (
	"flag"
	"gorbac/config"
	"gorbac/routes"
)

func main() {
	config.LoadEnv()
	db := config.InitialDB()

	flag.Parse()
	arg := flag.Arg(0)

	if arg != "" {
		config.InitCommands(db)
	} else {
		routes.WebRouter(db)
	}
}
