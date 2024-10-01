package main

import (
	"0byte/app"
	"0byte/config"
)

func main() {
	config.LoadConfigs()
	app.App()
}
