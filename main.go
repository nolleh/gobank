package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/nolleh/gobank/controllers"
	"github.com/nolleh/utils/dbcontext"
) 

func main() {
	appEnv := flag.String("app-env", os.Getenv("APP_ENV"), "app env")
	flag.Parse()
	var c Config
	if err := configutil.Read(*appEnv, &c); err != nil {
		panic(err)
	}

	fmt.Println(c)
	db, err := initDB(c.Database.Driver, c.Database.Connection)
	e := echo.New()

	controllers.BalanceController{}.Init(e.Group("/v1/balance"))

	e.Use(DbConext(db))
}

func initDB(driver, connection string) (*xorm.Engine, error) {
	db, err := xorm.NewEngine(driver, connection)
	if err != nil {
		return nil, err
	}

	db.Sync(new(models.balance))
	return db, nil
}

// config
type Config struct {
	Database struct {
		Driver string
		Connection string
	}
}