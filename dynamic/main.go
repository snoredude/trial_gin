package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    uint32      `json:"code"`
	Message uint32      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseError struct {
	Code    uint32 `json:"code"`
	Message uint32 `json:"message"`
}

// Handler
func IndexHandler(c *gin.Context) {
	// 前端HTML渲染(数据绑定，数据前后台传递)的方式:
	// {{ .title }}
	// {{ .message }}
	c.HTML(200, "index.html", gin.H{
		"title":   "主页",
		"message": "主页消息",
	})
}

func LoginHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title":   "登录页",
		"message": "登录页消息",
	})
}

// @title swagger使用例子
// @version 1.0
// @description swagger 入门使用例子

func main() {
	r := gin.Default()

	// 导入所有HTML文件，多级目录结构需要这样加载
	r.LoadHTMLGlob("html/*/*")
	// r.LoadHTMLFiles("templates/template1.html", "templates/template2.html") // 导入单个文件

	// 普通路由
	// r.GET("/index", IndexHandler) // 指定路由规则，访问路径 + 对应的Handler  http://localhost:8080/index
	// r.GET("/login", LoginHandler) // http://localhost:8080/login

	// 路由分组
	// 创建分组v
	v := r.Group("/")
	{
		v.GET("/index", IndexHandler)
		v.GET("/login", LoginHandler)
		v.GET("/check", connectCheck)
	}

	// 普通启动
	// r.Run()

	// 优雅启动（go1.8+）
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	go func() {
		// 监听请求
		// 使用ListenAndServe()监听8080端口，启动监听
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 优雅Shutdown（或重启）服务
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt) // syscall.SIGKILL
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
	}
	log.Println("Server exiting")
}

// @summary 服务连接校验 --> 接口简介
// @Description 服务初始连接测试 --> 接口描述
// @Accept json   --> 接收类型
// @Produce json  --> 返回类型
// Success 200 {object} Response --> 成功后返回数据结构
// Failure 400 {object} ResponseError --> 失败后返回数据结构
// Failure 404 {object} ResponseError
// Failure 500 {object} ResponseError
// @Router /check [get] --> 路由地址及请求方法
func connectCheck(c *gin.Context) {
	res := Response{Code: 1001, Message: 100, Data: "connect success !!!"}
	c.JSON(http.StatusOK, res)
}
