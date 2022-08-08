package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.StaticFS("/static", http.Dir("./static")) //加载所有静态资源
	router.LoadHTMLGlob("html/*/*")                  //加载所有HTML

	router.GET("/", defa)         // http://localhost:8080/
	router.GET("/signin", signin) // http://localhost:8080/login

	router.POST("/login", login)

	router.GET("/loginsucc", loginsucc)

	router.Run(":80")
}

func defa(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "首页",
	})
}

func signin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "登录页",
	})
}

func login(c *gin.Context) {
	username := c.DefaultPostForm("username", "bai")
	password := c.DefaultPostForm("password", "123")

	data := struct {
		username string
		password string
	}{
		username: username,
		password: password,
	}
	if username == "huahua" && password == "521" {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"data": data,
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": nil,
		})
	}
}

func loginsucc(c *gin.Context) {
	c.HTML(http.StatusOK, "loginsucc.html", gin.H{
		"username": "花花",
	})
}
