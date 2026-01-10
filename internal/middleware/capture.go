package middleware

import (
	"bytes"
	"io"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"webhook-receiver/internal/model"
)

func CaptureWebhook() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			provider := c.Param("provider")

			body, err := io.ReadAll(c.Request().Body)
			if err != nil {
				return echo.NewHTTPError(400, "Failed to read request body")
			}

			// restore body
			c.Request().Body = io.NopCloser(bytes.NewBuffer(body))

			headers := map[string]string{}
			for k, v := range c.Request().Header {
				if len(v) > 0 {
					headers[k] = v[0]
				}
			}

			event := &model.Event{
				ID:        uuid.NewString(),
				Provider:  provider,
				Headers:   headers,
				Body:      body,
				CreatedAt: time.Now(),
			}

			c.Set("event", event)

			return next(c)
		}
	}
}
