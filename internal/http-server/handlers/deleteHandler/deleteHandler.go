package deleteHandler

import (
	"net/http"
	"strconv"
	resp "web-service/internal/lib/api/response"
	"web-service/storage/postgres"

	"github.com/labstack/echo/v4"
)

func DeleteHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadGateway, resp.Error("Wrong id"))
	}

	if err := postgres.DB.Delete(&resp.Message{}, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, resp.Error("Could not delete the message"))
	}

	return c.JSON(http.StatusOK, resp.OK())
}
