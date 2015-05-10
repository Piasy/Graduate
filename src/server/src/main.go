package main

import (
	"github.com/astaxie/beego"

	_ "routers"
	"models"
)

func main() {
	defer models.CloseDB()
	beego.Run()
}
