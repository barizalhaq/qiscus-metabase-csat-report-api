package main

import (
	route "csat-report-webhook/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env")
}

func main() {
	r := gin.Default()

	r.Use(gin.Recovery())

	api := route.NewApi(r)
	api.Init()

	r.Run()
}
