package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"gobank/config"
	"gobank/echoMiddlewares"
	// "github.com/labstack/echo/middleware"
	"github.com/go-xorm/xorm"
	"gobank/api"
	"gobank/logger"
	"gobank/models"
	"log"
	"os"
)

func main() {
	appEnv := flag.String("app-env", os.Getenv("APP_ENV"), "app env")
	flag.Parse()

	configDir := "configFiles"
	config, err := config.GetConfig(*appEnv, configDir); if err != nil {
		panic(err)
	}
	fmt.Println(config)

	db, err := initDB(config.Database.Driver, config.Database.Connection)
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.Use(echoMiddlewares.InjectApiContext())
	e.Use(echoMiddlewares.InjectDbContext(db))
	e.Use(logger.ContextLogger())

	apiGroup := e.Group("/api")
	api.Route(apiGroup)

	if err := e.Start(":" + config.HttpPort); err != nil {
		log.Println(err)
	}
	fmt.Println("shutting down...")
}

func initDB(driver, connection string) (*xorm.Engine, error) {
	db, err := xorm.NewEngine(driver, connection)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)

	db.Sync(new(models.BalanceEntity))
	return db, nil
}