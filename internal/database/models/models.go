package models

import (
    "time"
    "encoding/json"

    "gorm.io/gorm"
)

// Database model for the exchange ticker
type ExchangeTicker struct {
    ID uint32 `gorm:"primaryKey"`
    Price float64
    Volume float64
    LastTrade float64
    CreatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
    SymbolID uint32
    Symbol Symbol
}

// Used when user requests exchange ticker from our server
func (et ExchangeTicker) MarshalJSON() ([]byte, error) {
    type etData struct {
        Price float64 `json:"price"`
        Volume float64 `json:"volume"`
        LastTrade float64 `json:"last_trade"`
    }
    data := etData{et.Price, et.Volume, et.LastTrade}
    entry := make(map[string]etData, 1)
    entry[et.Symbol.Name] = data
    return json.Marshal(entry)
}

// Database model for the symbol
type Symbol struct {
    ID uint32 `gorm:"primaryKey"`
    Name string `gorm:"unique"`
}
