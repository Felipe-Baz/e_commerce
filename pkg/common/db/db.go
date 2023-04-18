package db

import (
    "fmt"
    "log"

    "e_commerce/pkg/common/config"
    "e_commerce/pkg/common/models"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func Init(c *config.Config) *gorm.DB {
    // Configura url de coneção com banco de dados
    url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName)
    // Abre uma coneção com banco de dados
    db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }
    // Faz o migrations do product
    db.AutoMigrate(&models.Product{})
    // Retorna o banco de dados configurado
    return db
}