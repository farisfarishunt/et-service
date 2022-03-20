package handlers

import(
    "time"
    "errors"

    "gorm.io/gorm"

    "github.com/farisfarishunt/et-service/internal/database/models"
)

// Returns the exchange tickers that were last added to the database
// One for each symbol
func GetLastExchangeTickers(db *gorm.DB) ([]models.ExchangeTicker, error) {
    var exchangeTickers []models.ExchangeTicker
    subQuery := db.Model(&models.ExchangeTicker{}).Select("symbol_id, max(created_at) as last_created_at").Group("symbol_id")
    err := db.Preload("Symbol").Joins("join (?) sd on created_at = sd.last_created_at and exchange_tickers.symbol_id = sd.symbol_id", subQuery).Find(&exchangeTickers).Error

    return exchangeTickers, err
}

// Adds exchange tickers to the database
func AddExchangeTicker(db *gorm.DB, symbolName string, price float64, volume float64, lastTrade float64) error {
    tx := db.Begin()
    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        }
    } ()

    var symbol models.Symbol
    err := tx.Where("name = ?", symbolName).First(&symbol).Error
    // We're adding the symbol if it's not stored yet in the database
    if errors.Is(err, gorm.ErrRecordNotFound) {
        symbol.Name = symbolName;
        if err := tx.Create(&symbol).Error; err != nil {
            return err
        }
    } else if err != nil {
        return err
    }
    exchangeTicker := models.ExchangeTicker{
        Price: price,
        Volume: volume,
        LastTrade: lastTrade,
        CreatedAt: time.Now(),
        SymbolID: symbol.ID,
    }
    if err := tx.Create(&exchangeTicker).Error; err != nil {
        return err
    }
    
    if err := tx.Commit().Error; err != nil {
        return err
    }
    
    return nil
}
