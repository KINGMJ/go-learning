package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	bindUriDemo()
}

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Age      int       `form:"age"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func bindFormDemo() {
	r := gin.Default()
	var person Person
	r.GET("/testing", func(ctx *gin.Context) {
		// 绑定到 query
		if err := ctx.Bind(&person); err != nil {
			ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, person)
	})
	r.Run()
}

type Person1 struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func bindUriDemo() {
	r := gin.Default()
	var person Person1
	r.GET("/:name/:id", func(ctx *gin.Context) {
		if err := ctx.ShouldBindUri(&person); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"name": person.Name, "uuid": person.ID})
	})
	r.Run()
}
