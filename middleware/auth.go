package middleware

import (
    "net/http"
    "github.com/labstack/echo/v4"
)

func IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        cookie, err := c.Cookie("session_id")
        if err != nil {
            return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
        }

        if cookie.Value == "" {
            return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
        }
        return next(c)
    }
}
