package service

import (
    "time"

    "github.com/go-co-op/gocron"
    "github.com/gin-gonic/gin"

    "github.com/farisfarishunt/et-service/internal/config"
    "github.com/farisfarishunt/et-service/internal/database/connection"
    "github.com/farisfarishunt/et-service/internal/cron/spawner"
    "github.com/farisfarishunt/et-service/internal/cron/jobs"
    "github.com/farisfarishunt/et-service/internal/server/router"
)

// Performs server systems initialization and returns it's parts
func NewService() (server *gin.Engine, cfg *config.Config, err error) {
    cfg, err = config.NewConfig()
    if err != nil {
        return
    }

    db, err := database.NewConnection(cfg.DbHost, cfg.DbUser, cfg.DbPassword, cfg.DbPort, cfg.DbName)
    if err != nil {
        return
    }

    // Init and then start cron jobs (job will repeat at regular intervals)
    scheduler := gocron.NewScheduler(time.UTC)
    cron.Spawn(
        scheduler,
        grabber.GrabExchangeTickers(db, cfg.ExchangeTickerUrl, scheduler.CronWithSeconds(cfg.CronGrabInterval)),
    )

    server = router.NewRouter(db)

    return
}
