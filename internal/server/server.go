package server

import (
	"log"
	"webhook-receiver/internal/handler"
	"webhook-receiver/internal/middleware"
	"webhook-receiver/internal/model"
	"webhook-receiver/internal/store"

	"github.com/labstack/echo/v4"
)

func Start(addr string, store *store.MemoryStore) {
	e := echo.New()

	e.Use(middleware.CaptureWebhook())

	// Webhook受信エンドポイント
	e.POST("/hooks/:provider", func(c echo.Context) error {
		event := c.Get("event").(*model.Event)
		store.SaveEvent(event)
		return handler.ReceiveWebhook(c)
	})

	e.GET("/events", func(c echo.Context) error {
		return c.JSON(200, store.GetAllEvents())
	})

	e.GET("/events/:id", func(c echo.Context) error {
		id := c.Param("id")
		event, ok := store.GetEventByID(id)
		if !ok {
			return echo.NewHTTPError(404, "Event not found")
		}
		return c.JSON(200, event)
	})

	log.Printf("🚀 Listening on %s\n", addr)
	if err := e.Start(addr); err != nil {
		log.Fatal(err)
	}

	e.Start(addr)
}
