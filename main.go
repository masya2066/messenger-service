package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"pager-service/routes"
)

func main() {
	r := gin.Default()
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	routes.Calls(r)

	err = r.Run(":1090")
	if err != nil {
		return
	}
}
