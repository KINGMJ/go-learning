package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	routeGroupDemo()
}

// 上传单个文件
func uploadFileDemo() {
	r := gin.Default()
	// 限制上传文件大小，默认为 32M
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
		// 保存文件到当前目录
		c.SaveUploadedFile(file, file.Filename)
		c.JSON(http.StatusOK, file)
	})
	r.Run()
}

// 上传多个文件
func uploadMultiFileDemo() {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
		}
		// 获取所有文件
		files := form.File["files"]
		// 遍历每个文件
		for _, file := range files {
			if err := c.SaveUploadedFile(file, file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
				return
			}
		}
		c.JSON(http.StatusOK, files)
	})
	r.Run()
}

func routeGroupDemo() {
	r := gin.Default()
	// 路由组 v1，处理版本 v1 的请求
	// http://localhost:8080/v1/login
	v1 := r.Group("/v1")
	{
		v1.GET("/login", func(c *gin.Context) {
			c.String(http.StatusOK, "login v1")
		})
		v1.GET("/logout", func(c *gin.Context) {
			c.String(http.StatusOK, "logout v1")
		})
	}

	// 路由组 v2，处理版本 v2 的请求
	v2 := r.Group("/v2")
	{
		v2.GET("/login", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "login v2")
		})
		v2.GET("/logout", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "logout v2")
		})
	}
	// 非路由组请求
	r.GET("/login", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "login not version")
	})

	r.Run()
}
