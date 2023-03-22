package main

import (
	"ginblog/models"
	"ginblog/routers"
)

func main() {
	models.InitDb()
	routers.InitRouter()
}
