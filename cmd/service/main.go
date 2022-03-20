package main

import (
    "fmt"

    "github.com/farisfarishunt/et-service/internal/service"
    "github.com/farisfarishunt/et-service/internal/logger"
)

func main() {
    db, cronDb, server, cfg, err := service.NewService()
    if db != nil {
        defer func() {
            if db, _ := db.DB(); db != nil {
                db.Close()
            }
        } ()
    }
    if cronDb != nil {
        defer func() {
            if cronDb, _ := cronDb.DB(); cronDb != nil {
                cronDb.Close()
            }
        } ()
    }
    if err != nil {
        logger.Log(err)
        return
    }

    if err = server.Run(fmt.Sprintf(":%d", cfg.ServerPort)); err != nil {
        logger.Log(err)
    }
}
