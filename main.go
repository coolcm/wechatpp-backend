package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sjtucsn/wechatpp-backend/controller"
	"github.com/sjtucsn/wechatpp-backend/model"
)

func setupRouter() *gin.Engine {

	r := gin.Default()

	// Get user value
	r.GET("/user/create", controller.HandleCreateUser)
	r.GET("/user/query", controller.HandleQueryUser)

	return r
}

func main() {
	defer func() {
		model.Db.Close()
		println("hahaha goodbye")
	}()
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}