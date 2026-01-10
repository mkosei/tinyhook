package main

import (
	"webhook-receiver/internal/handler"
	"webhook-receiver/internal/middleware"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Middlewareを登録
	e.Use(middleware.CaptureWebhook())

	// Webhook受信エンドポイント
	e.POST(
		"/hooks/:provider",
		handler.ReceiveWebhook,
		middleware.CaptureWebhook(),
	)

	e.Logger.Fatal(e.Start(":8080"))
}
