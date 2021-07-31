package main

import (
	"buhang/bootstrap"
	"buhang/routes"
	"github.com/gobuffalo/packr"
)

func main() {
	bootstrap.Init()
	routes.Init()
	routes.R.StaticFS("/manager", packr.NewBox("./public/manager"))
	routes.R.Run()
}
