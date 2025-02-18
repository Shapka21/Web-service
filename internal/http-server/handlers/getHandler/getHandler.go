package getHandler

import (
	"net/http"
	resp "web-service/internal/lib/api/response"
	"web-service/storage/postgres"

	"github.com/labstack/echo/v4"
)

func GetHandler(c echo.Context) error {
	var messages []resp.Message

	if err := postgres.DB.Find(&messages).Error; err != nil {
		return c.JSON(http.StatusBadRequest, resp.Error("Could not find the message"))
	}

	return c.JSON(http.StatusOK, &messages)
}
