package service

import (
    "time"

    "gorm.io/gorm"
    "github.com/go-co-op/gocron"
    "github.com/gin-gonic/gin"

    "github.com/farisfarishunt/et-service/internal/config"
    "github.com/farisfarishunt/et-service/internal/database/connection"
    "github.com/farisfarishunt/et-service/internal/cron/spawner"
    "github.com/farisfarishunt/et-service/internal/cron/jobs"
    "github.com/farisfarishunt/et-service/internal/server/router"
)

// Performs server systems initialization and returns it's parts
func NewService() (db *gorm.DB, cronDb *gorm.DB, server *gin.Engine, cfg *config.Config, err error) {
    cfg, err = config.NewConfig()
    if err != nil {
        return
    }

    db, err = database.NewConnection(cfg.DbHost, cfg.DbUser, cfg.DbPassword, cfg.DbPort, cfg.DbName)
    if err != nil {
        return
    }
    
    cronDb, err = database.NewConnection(cfg.DbHost, cfg.DbUser, cfg.DbPassword, cfg.DbPort, cfg.DbName)
    if err != nil {
        return
    }

    // Init and then start cron jobs (job will repeat at regular intervals)
    scheduler := gocron.NewScheduler(time.UTC)
    cron.Spawn(
        scheduler,
        grabber.GrabExchangeTickers(cronDb, cfg.ExchangeTickerUrl, scheduler.CronWithSeconds(cfg.CronGrabInterval)),
    )

    server = router.NewRouter(db)

    return
}
