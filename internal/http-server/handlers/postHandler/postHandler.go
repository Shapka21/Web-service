package postHandler

import (
	"net/http"
	resp "web-service/internal/lib/api/response"
	"web-service/storage/postgres"

	"github.com/labstack/echo/v4"
)

func PostHandler(c echo.Context) error {
	var message resp.Message

	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, resp.Error("Could not add the message"))
	}

	if err := postgres.DB.Create(&message).Error; err != nil {
		return c.JSON(http.StatusBadRequest, resp.Error("Could not create the message"))
	}

	return c.JSON(http.StatusOK, resp.OK())
}
