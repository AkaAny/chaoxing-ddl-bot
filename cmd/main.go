package main

import (
	"ddl-bot/internal"
	"ddl-bot/internal/config"
)

func main() {
	conf := config.Unmarshall("config/conf.toml")
	app := internal.Create(conf)
	app.Run()
}
