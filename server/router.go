package server

import "github.com/labstack/echo/v4"

func ProvideRouter(e *echo.Echo) {
	h := handler.NewHandler()
}
