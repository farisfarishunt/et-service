package handlers

import (
    "net/http"

    "gorm.io/gorm"
    "github.com/gin-gonic/gin"

    db_handlers "github.com/farisfarishunt/et-service/internal/database/handlers"
    "github.com/farisfarishunt/et-service/internal/logger"
)

type ApiHandler struct {
    Db *gorm.DB
}

// Handler fired when user retrieves exchange tickers from our database
func (h *ApiHandler) GetExchangeTickers(ctx *gin.Context) {
    db := h.Db.WithContext(ctx)
    exchangeTickers, err := db_handlers.GetLastExchangeTickers(db)
    if err != nil {
        logger.Log(err)
        const errorMessage = "Internal error. Can't get exchange tickers."
        ctx.JSON(http.StatusInternalServerError, errorMessage)
        return
    }

    ctx.JSON(http.StatusOK, exchangeTickers)
}
