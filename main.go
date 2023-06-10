package main

import (
	"ginblog/models"
	"ginblog/routers"
)

func main() {
	models.InitDb()
	models.InitRedis()
	routers.InitRouter()
}
