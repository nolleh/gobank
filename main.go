package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	// "github.com/labstack/echo/middleware"
	"github.com/go-xorm/xorm"
	"gobank/api"
	"gobank/logger"
	"gobank/models"
	"gobank/utils"
	"log"
	"os"
)

func main() {
	appEnv := flag.String("app-env", os.Getenv("APP_ENV"), "app env")
	flag.Parse()

	configDir := "config"
	config, err := utils.GetConfig(*appEnv, configDir); if err != nil {
		panic(err)
	}
	fmt.Println(config)

	db, err := initDB(config.Database.Driver, config.Database.Connection)
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.Use(utils.DbContext(db))
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

	db.Sync(new(models.Balance))
	return db, nil
}