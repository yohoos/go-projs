package main

import (
	"go-projs/portfolio/controllers"
)

func main() {
	a := controllers.App{}
	a.Initialize("yohoos", "magicdust50", "testdb")
	a.Run(":8080")
}
