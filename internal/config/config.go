package config

import (
    "os"

    "github.com/asaskevich/govalidator"
    "github.com/namsral/flag"
)

// Represents config variables
type Config struct {
    DbHost string `valid:"required"`
    DbName string `valid:"required"`
    DbUser string `valid:"required"`
    DbPassword string `valid:"required"`
    DbPort int `valid:"port,required"`
    ServerPort int `valid:"port,required"`
    CronGrabInterval string `valid:"required"`
    ExchangeTickerUrl string `valid:"url,required"`
}

// Validates config variables. If at least one of them doesn't fit to the requirements then error will be not empty
func (config Config) validate() error {
    if _, err := govalidator.ValidateStruct(config); err != nil {
        return err
    }

    return nil
}

// Returns new config and validation error (if any). Config variables are passed by env variables or argument parameters
func NewConfig() (*Config, error) {
    config := Config{}
    fs := flag.NewFlagSetWithEnvPrefix(os.Args[0], "BLOCKCHAIN", flag.ExitOnError)
    fs.StringVar(&config.DbHost, "db_host", "", "Database host")
    fs.StringVar(&config.DbName, "db_name", "", "Database name")
    fs.StringVar(&config.DbUser, "db_user", "", "Database user")
    fs.StringVar(&config.DbPassword, "db_password", "", "Database password")
    fs.IntVar(&config.DbPort, "db_port_internal", -1, "Database port")
    fs.IntVar(&config.ServerPort, "grabber_service_port_internal", -1, "Server port")
    fs.StringVar(&config.CronGrabInterval, "grabber_service_cron_grab_interval", "", "Cron-pattern interval for grabbing exchange tickers")
    fs.StringVar(&config.ExchangeTickerUrl, "grabber_service_exchange_ticker_url", "", "Url to grab exchange tickers")
    fs.Parse(os.Args[1:])
    return &config, config.validate()
}
