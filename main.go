package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

// Note don't read this code base as it will not feel good about its structure
func main() {
	svr := echo.New()
	dto := NewDTO(data{})
	handler := NewHandler(dto)

	svr.GET("/version", handler.version)
	svr.GET("/", handler.list)
	svr.GET("/:key", handler.get)
	svr.POST("/", handler.post)
	svr.DELETE("/:key", handler.delete)

	err := svr.Start(fmt.Sprintf(":8080")); if err != nil {
		log.Fatal(err)
	}
}

