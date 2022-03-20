package router

import (
    "gorm.io/gorm"
    "github.com/gin-gonic/gin"

    "github.com/farisfarishunt/et-service/internal/server/handlers"
)

// Returns router that can be used further to run the server
func NewRouter(db *gorm.DB) *gin.Engine {
    router := gin.Default()
    apiHandler := &handlers.ApiHandler {
        Db: db,
    }
    
    router.GET("", apiHandler.GetExchangeTickers)

    return router
}
