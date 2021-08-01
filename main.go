package main

import (
	"buhang/bootstrap"
	"buhang/config"
	"buhang/routes"
)

func main() {
	config.Init()
	bootstrap.Init()
	routes.Init()
}
