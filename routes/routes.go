package routes

import (
    "github.com/labstack/echo/v4"
    "echoAPI/controllers"
    //"echoAPI/middleware"
)

func InitRoutes(e *echo.Echo) {
    // Rute untuk otentikasi
    // e.POST("/login", controllers.Login)
    // e.POST("/logout", controllers.Logout)

    // Rute API buku
    r := e.Group("/api")
    //r.Use(middleware.IsAuthenticated)
    r.GET("/books", controllers.GetBooks)
    r.GET("/books/:id", controllers.GetBookByID)
    r.POST("/books", controllers.CreateBook)
    r.PUT("/books/:id", controllers.UpdateBook)
    r.DELETE("/books/:id", controllers.DeleteBook)
}
