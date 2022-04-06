package handlers

import (
    "testing"
    "database/sql"
    "regexp"
    "time"

    "github.com/stretchr/testify/suite"
    "github.com/stretchr/testify/require"
    "github.com/DATA-DOG/go-sqlmock"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"

    "github.com/farisfarishunt/et-service/internal/database/models"
)

type DataRows struct {
    Rows *sqlmock.Rows
}

type TestDataType struct {
    Expec []models.ExchangeTicker
    SingleExchangeTicker DataRows
    SingleSymbol DataRows
    ExchangeTickers DataRows
    Symbols DataRows
}

type DbHandlersSuite struct {
    suite.Suite
    TestData TestDataType
    Db *gorm.DB
    Mock sqlmock.Sqlmock
}

func (suite *DbHandlersSuite) InitTestData() {
    suite.TestData.Expec = make([]models.ExchangeTicker, 0, 2)
    suite.TestData.Expec = append(suite.TestData.Expec, models.ExchangeTicker{
        ID: 1,
        Price: 30.0,
        Volume: 45.4,
        LastTrade: 31.3,
        CreatedAt: time.Now(),
        DeletedAt: gorm.DeletedAt{Valid: false},
        SymbolID: 1,
        Symbol: models.Symbol{ID: 1, Name: "BTC"},
    })
    suite.TestData.Expec = append(suite.TestData.Expec, models.ExchangeTicker{
        ID: 2,
        Price: 1.0,
        Volume: 3.0,
        LastTrade: 2.0,
        CreatedAt: time.Now(),
        DeletedAt: gorm.DeletedAt{Valid: false},
        SymbolID: 2,
        Symbol: models.Symbol{ID: 2, Name: "ETH"},
    })
    suite.TestData.ExchangeTickers.Rows = sqlmock.NewRows([]string{"id", "price", "volume", "last_trade", "created_at", "deleted_at", "symbol_id"})
    suite.TestData.Symbols.Rows = sqlmock.NewRows([]string{"id", "name"})
    for _, et := range suite.TestData.Expec {
        suite.TestData.ExchangeTickers.Rows = suite.TestData.ExchangeTickers.Rows.
                  AddRow(et.ID, et.Price, et.Volume, et.LastTrade, et.CreatedAt, et.DeletedAt, et.Symbol.ID)
        suite.TestData.Symbols.Rows = suite.TestData.Symbols.Rows.
                      AddRow(et.Symbol.ID, et.Symbol.Name)
    }

    suite.TestData.SingleExchangeTicker.Rows = sqlmock.NewRows([]string{"id", "price", "volume", "last_trade", "created_at", "deleted_at", "symbol_id"}).AddRow(suite.TestData.Expec[0].ID, suite.TestData.Expec[0].Price, suite.TestData.Expec[0].Volume, suite.TestData.Expec[0].LastTrade, suite.TestData.Expec[0].CreatedAt, suite.TestData.Expec[0].DeletedAt, suite.TestData.Expec[0].Symbol.ID)

    suite.TestData.SingleSymbol.Rows = sqlmock.NewRows([]string{"id", "name"}).AddRow(suite.TestData.Expec[0].Symbol.ID, suite.TestData.Expec[0].Symbol.Name)
}

func (suite *DbHandlersSuite) SetupSuite() {
    var (
        db *sql.DB
        err error
    )

    db, suite.Mock, err = sqlmock.New()
    require.NoError(suite.T(), err)
    suite.Mock.MatchExpectationsInOrder(false)

    suite.Db, err = gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
    require.NoError(suite.T(), err)

    suite.InitTestData()
}

func (suite *DbHandlersSuite) TestGetLastExchangeTickers() {
    sqlQueryEts := regexp.QuoteMeta(`SELECT "exchange_tickers"."id","exchange_tickers"."price","exchange_tickers"."volume","exchange_tickers"."last_trade","exchange_tickers"."created_at","exchange_tickers"."deleted_at","exchange_tickers"."symbol_id" FROM "exchange_tickers" join (SELECT symbol_id, max(created_at) as last_created_at FROM "exchange_tickers" WHERE "exchange_tickers"."deleted_at" IS NULL GROUP BY "symbol_id") sd on created_at = sd.last_created_at and exchange_tickers.symbol_id = sd.symbol_id WHERE "exchange_tickers"."deleted_at" IS NULL`)
    sqlQuerySymbols := regexp.QuoteMeta(`SELECT * FROM "symbols" WHERE "symbols"."id" IN ($1,$2)`)

    suite.Mock.ExpectQuery(sqlQueryEts).WillReturnRows(suite.TestData.ExchangeTickers.Rows)
    suite.Mock.ExpectQuery(sqlQuerySymbols).WithArgs(suite.TestData.Expec[0].Symbol.ID, suite.TestData.Expec[1].Symbol.ID).WillReturnRows(suite.TestData.Symbols.Rows)

    ets, err := GetLastExchangeTickers(suite.Db)
    require.NoError(suite.T(), err)
    require.Equal(suite.T(), suite.TestData.Expec, ets)

    err = suite.Mock.ExpectationsWereMet()
    require.NoError(suite.T(), err)
}

func TestDbHandlerSuite(t *testing.T) {
    suite.Run(t, new(DbHandlersSuite))
}
