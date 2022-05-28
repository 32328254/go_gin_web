//
// package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func api(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"status": "ok",
// 	})
// }
// func main() {
// 	//创建一个默认的gin引擎
// 	r := gin.Default()
// 	//指定用户使用get方式请求资源
// 	r.GET("/hello", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "Hello world",
// 		})
// 	})
// 	r.GET("/api", api)

// 	//启动http服务，默认是8080端口
// 	r.Run("0.0.0.0:9090")
// }

//