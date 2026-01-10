package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ReceiveWebhook(c echo.Context) error {
	raw := c.Get("rawBody").([]byte)
	provider := c.Param("provider")

	// ここでrawとproviderを使って処理を行う
	return c.JSON(http.StatusOK, map[string]interface{}{
		"provider": provider,
		"body":     string(raw),
	})
}
