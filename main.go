package main

import (
	"log"
	"github.com/wangjiaming0909/GoP/dbService"
	"github.com/wangjiaming0909/GoP/userservice"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	dbCf := dbService.NewDBServiceConfig("mysql", "root", "123456", "sys", "127.0.0.1", 3306)
	dbS := dbService.NewDBService(&dbCf)
	err := dbS.Start()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer dbS.Stop()

	r := gin.Default()

	r.GET("/ping", Pong)
	r.POST("/user", userservice.UserRegisterCallBack)
	r.Run(":10001")
	r.Static("/", "./1.jpg")

}

func Pong(c *gin.Context) {
	name := c.Param("name")
	dbService.SqlDaoService()
	c.String(http.StatusOK, "Hello %s", name)
}
