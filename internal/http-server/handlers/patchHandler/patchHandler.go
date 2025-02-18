package patchHandler

import (
	"net/http"
	"strconv"
	resp "web-service/internal/lib/api/response"
	"web-service/storage/postgres"

	"github.com/labstack/echo/v4"
)

func PatchHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resp.Error("Wrong id"))
	}

	var updatedMessage resp.Message
	if err = c.Bind(&updatedMessage); err != nil {
		return c.JSON(http.StatusBadGateway, resp.Error("Invalid input"))
	}

	if err := postgres.DB.Model(&resp.Message{}).Where("id = ?", id).Update("text", updatedMessage.Text).Error; err != nil {
		return c.JSON(http.StatusBadRequest, resp.Error("Could not update the message"))
	}

	return c.JSON(http.StatusOK, resp.OK())
}
