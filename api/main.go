package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	UserName string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Pssword  string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
}

func main() {
	/*
		// hello world
		// 1.创建路由
		router := gin.Default()

		// 2.绑定路由规则，执行的函数
		// gin.Context，封装了request和response
		router.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "hello world!")
		})

		// 3.监听端口，默认在8080
		// Run("里面不指定端口号默认为8080")
		// 访问: localhost:8000
		router.Run(":8000")
	*/

	/*
		router := gin.Default()
		router.GET("/user", func(c *gin.Context) {
			// 不传参, 使用默认值 http://localhost:8080/user 才会打印出来默认的值 hello xiaoming
			// 传参, 不使用默认值 http://localhost:8080/user?name=zhangsan hello zhangsan
			name := c.DefaultQuery("name", "xiaoming")
			c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
		})
		router.Run()
	*/

	/*
		// 表单传输为post请求，http常见的传输格式为四种：
		// application/json
		// application/x-www-form-urlencoded
		// application/xml
		// multipart/form-data
		// 表单参数可以通过PostForm()方法获取, ，该方法默认解析的是x-www-form-urlencoded或from-data格式的参数
		r := gin.Default()
		r.POST("/form", func(c *gin.Context) {
			types := c.DefaultPostForm("type", "post")
			username := c.PostForm("username")
			password := c.PostForm("userpassword")
			c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
		})
		r.Run()
	*/

	/*
		// 路由分组
		r := gin.Default()
		v1 := r.Group("/v1")
		{
			v1.GET("/login", login)
			v1.GET("submit", submit)
		}
		v2 := r.Group("/v2")
		{
			v2.POST("/login", login)
			v2.POST("/submit", submit)
		}
		r.Run(":8000")
		// get: curl http://localhost:8000/v1/login?name=zhangsan  hello zhangsan
		// post: curl http://localhost:8000/v1/submit -X POST   page not found
		// post: curl http://localhost:8000/v2/submit -X POST   hello lily
	*/

	/*
		// 404
		r := gin.Default()
		r.GET("/user", func(c *gin.Context) {
			name := c.DefaultQuery("name", "zhangsan")
			c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
		})
		r.NoRoute(func(c *gin.Context) {
			c.String(http.StatusNotFound, "404 not found")
		})
		r.Run()
	*/

	// json
	r := gin.Default()
	r.POST("loginJSON", func(c *gin.Context) {
		var loginfo Login
		if err := c.ShouldBindJSON(&loginfo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 判断用户名密码是否正确
		if loginfo.UserName != "root" || loginfo.Pssword != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	r.Run(":8000")
	// curl http://localhost:8000/loginJSON -H 'content-type:application/json' -d {"user": "root", "password": "admin"} -X POST   {"error":"invalid character 'u' looking for beginning of object key string"}
	// curl http://localhost:8000/loginJSON -H 'content-type:application/json' -d "{\"user\": \"root\", \"password\": \"admin\"}" -X POST   {"status":"200"}
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "lily")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

// https://www.kancloud.cn/shuangdeyu/gin_book/949414
// http://baike.eepw.com.cn/baike/show/word/%E8%80%A6%E5%90%88%E6%80%A7
