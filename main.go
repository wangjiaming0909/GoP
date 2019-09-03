package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wangjiaming0909/GoP/database"
	"github.com/wangjiaming0909/GoP/userservice"
)

func main() {
	r := gin.Default()

	r.GET("/ping", Pong)
	r.POST("/user", userservice.UserRegisterCallBack)
	r.Run(":10001")
	r.Static("/", "./1.jpg")

}

func Pong(c *gin.Context) {
	name := c.Param("name")
	database.SqlDaoService()
	c.String(http.StatusOK, "Hello %s", name)
}
