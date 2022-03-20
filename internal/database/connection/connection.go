package database

import (
    "fmt"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"

    "github.com/farisfarishunt/et-service/internal/database/models"
)

// Establishes database connection and returns pointer to db object
func NewConnection(host string, user string, password string, port int, dbName string) (*gorm.DB, error) {
    dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, password, host, port, dbName)

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        return nil, err
    }

    db.AutoMigrate(&models.Symbol{}, &models.ExchangeTicker{})

    return db, nil
} 
