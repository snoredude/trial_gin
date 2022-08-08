package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// router是一个gin Engine，是gin的核心，默认带有 Logger 和 Recovery 两个中间件
	router := gin.Default()

	// 静态资源加载，css/js/images/fonts/icons...
	// StaticFile 是加载单个文件，而 StaticFS 是加载一个完整的目录资源
	router.StaticFS("/images", http.Dir("./images"))  // 访问: http://localhost:8080/image/ram.jpg 实际访问的服务器目录是: ./images/ram.jpg
	router.StaticFile("/icon", "./icon/mini.ico")     // 访问: http://localhost:8080/icon 实际访问的服务器目录是: ./icon/mini.ico
	router.StaticFile("/together", "./images/sd.jpg") //
	// Listen and serve on 0.0.0.0:8080
	// 编译运行程序，静态站点就可以正常访问了
	// 每次请求响应都会在服务端有日志产生，包括响应时间，加载资源名称，响应状态值等等
	router.Run(":80")
}
