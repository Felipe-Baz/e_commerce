package main

import (
    "log"

    "e_commerce/pkg/common/config"
    "e_commerce/pkg/common/db"
    "e_commerce/pkg/products"

    "github.com/gofiber/fiber/v2"
)

func main() {
    configs, err := config.LoadConfig()

    if err != nil {
        log.Fatalln("Failed to load configs", err)
    }

    database := db.Init(&configs)
    app := fiber.New()

    products.RegisterRoutes(app, database)

    app.Listen(configs.Port)
}