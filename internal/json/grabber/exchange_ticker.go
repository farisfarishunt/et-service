package jsn

// Used when data retrieved from exchange ticker url
type ExchangeTicker struct {
    Symbol string `json:"symbol" valid:"required~Symbol is blank"`
    Price float64 `json:"price_24h"`
    Volume float64 `json:"volume_24h"`
    LastTrade float64 `json:"last_trade_price"`
}
