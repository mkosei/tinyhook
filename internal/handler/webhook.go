package handler

import (
	"net/http"

	"webhook-receiver/internal/model"

	"github.com/labstack/echo/v4"
)

func ReceiveWebhook(c echo.Context) error {
	event := c.Get("event").(*model.Event)

	return c.JSON(http.StatusOK, map[string]any{
		"status":   "received",
		"id":       event.ID,
		"provider": event.Provider,
		"size":     len(event.Body),
	})
}
