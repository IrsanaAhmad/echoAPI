package main

import (
    "echoAPI/config"
    "github.com/labstack/echo/v4"
    "echoAPI/routes"
)

func main() {
    e := echo.New()

    config.InitDB()
    defer config.DB.Close()
    routes.InitRoutes(e)
    e.Logger.Fatal(e.Start(":1323"))
}
