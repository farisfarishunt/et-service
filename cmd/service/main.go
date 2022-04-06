package main

import (
    "fmt"

    "github.com/farisfarishunt/et-service/internal/service"
    "github.com/farisfarishunt/et-service/internal/logger"
)

func main() {
    server, cfg, err := service.NewService()
    if err != nil {
        logger.Log(err)
        return
    }

    if err = server.Run(fmt.Sprintf(":%d", cfg.ServerPort)); err != nil {
        logger.Log(err)
    }
}
