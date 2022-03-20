package grabber 

import(
    "net/http"
    "encoding/json"
    "time"

    "gorm.io/gorm"
    "github.com/go-co-op/gocron"
    "github.com/asaskevich/govalidator"

    "github.com/farisfarishunt/et-service/internal/cron/spawner"
    "github.com/farisfarishunt/et-service/internal/logger"
    "github.com/farisfarishunt/et-service/internal/json/grabber"
    db_handlers "github.com/farisfarishunt/et-service/internal/database/handlers"
)

// Retrieves exchange tickers from the url
func retrieveJson(exchangeTickerUrl string) ([]jsn.ExchangeTicker, error) {
    r, err := (&http.Client{Timeout: 10 * time.Second}).Get(exchangeTickerUrl)
    if err != nil {
        return nil, err
    }
    defer r.Body.Close()
    
    var exchangeTickers []jsn.ExchangeTicker
    if err := json.NewDecoder(r.Body).Decode(&exchangeTickers); err != nil {
        return nil, err
    }

    return exchangeTickers, nil
}

// Retrieves exchange tickers from the url and puts it into the database
// Returns function to be scheduled with cron scheduler
func grabExchangeTickers(db *gorm.DB, exchangeTickerUrl string) func() {
    return func() {
        exchangeTickerArr, err := retrieveJson(exchangeTickerUrl)
        if err != nil {
            logger.Log(err)
            return
        }
        for _, et := range exchangeTickerArr {
            if _, err := govalidator.ValidateStruct(&et); err != nil {
                logger.Log(err)
                continue
            }
            if err := db_handlers.AddExchangeTicker(db, et.Symbol, et.Price, et.Volume, et.LastTrade); err != nil {
                logger.Log(err)
                continue
            }
        }
    }
}

// Returns the job wrapper to be scheduled with cron scheduler
func GrabExchangeTickers(db *gorm.DB, exchangeTickerUrl string, interval *gocron.Scheduler) cron.Job {
    return cron.Job {
        JobFunc: grabExchangeTickers(db, exchangeTickerUrl),
        Interval: interval,
    }
}
