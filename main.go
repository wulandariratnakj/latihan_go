package main

import (
	"prakerja2/config"
	"prakerja2/route"
)

func main() {
	config.InitDB()
	e := route.InitRoute()
	e.Logger.Fatal(e.Start(":8080"))
}
