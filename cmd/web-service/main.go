package main

import (
	delete "web-service/internal/http-server/handlers/deleteHandler"
	get "web-service/internal/http-server/handlers/getHandler"
	patch "web-service/internal/http-server/handlers/patchHandler"
	post "web-service/internal/http-server/handlers/postHandler"
	"web-service/storage/postgres"

	"github.com/labstack/echo/v4"
)

func main() {
	postgres.InitDB()

	e := echo.New()

	e.GET("/messages", get.GetHandler)
	e.POST("/messages", post.PostHandler)
	e.PATCH("/messages/:id", patch.PatchHandler)
	e.DELETE("messages/:id", delete.DeleteHandler)

	e.Start(":8080")
}
