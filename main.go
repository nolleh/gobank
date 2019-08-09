package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"

	"github.com/labstack/echo"
	// "github.com/labstack/echo/middleware"
	"github.com/nolleh/gobank/controllers"
	"github.com/nolleh/gobank/utils"
	"github.com/nolleh/gobank/models"
) 

func main() {
	appEnv := flag.String("app-env", os.Getenv("APP_ENV"), "app env")
	flag.Parse()
	var c Config
	if err := utils.ReadConfig(*appEnv, &c); err != nil {
		panic(err)
	}

	fmt.Println(c)
	db, err := initDB(c.Database.Driver, c.Database.Connection)
	if err != nil {
		panic(err)
	}
	e := echo.New()

	controllers.BalanceController{}.Init(e.Group("/v1/balance"))

	if err := e.Start(":" + c.HttpPort); err != nil {
		log.Println(err)
	}
	e.Use(utils.DbContext(db))
	fmt.Println("shutting down...")
}

func initDB(driver, connection string) (*xorm.Engine, error) {
	db, err := xorm.NewEngine(driver, connection)
	if err != nil {
		return nil, err
	}

	db.Sync(new(models.Balance))
	return db, nil
}

// config
type Config struct {
	Database struct {
		Driver string
		Connection string
	}
	HttpPort string
}