package middleware

import (
	"bytes"
	"io"

	"github.com/labstack/echo/v4"
)

func CaptureBodyMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			// Bodyを読む
			body, err := io.ReadAll(req.Body)
			if err != nil {
				return err
			}
			//Bodyを戻す
			req.Body = io.NopCloser(bytes.NewBuffer(body))
			//Contextに保存
			c.Set("rawBody", body)

			return next(c)
		}
	}
}
