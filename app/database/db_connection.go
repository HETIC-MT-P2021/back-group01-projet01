package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/caarlos0/env/v6"
	_ "github.com/go-sql-driver/mysql"
	cLog "image_gallery/logger"
	"time"
)

//DbConn stores the connexion to the database
var (
	DbConn *sql.DB
)

type Config struct {
	DbHost     string `env:"DB_HOST"`
	DbName     string `env:"MYSQL_DATABASE"`
	DbUser     string `env:"MYSQL_USER"`
	DbPassword string `env:"MYSQL_PASSWORD"`
	DbConn     *sql.DB
}

// Connect connection to database
func Connect() error {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}
	dsn := cfg.DbUser + ":" + cfg.DbPassword + "@" + cfg.DbHost + "/" + cfg.
		DbName + "?parseTime=true&charset=utf8"

	logger := cLog.GetLogger()

	logger.Infof("DSN: %s", dsn)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return err
	}

	var dbErr error
	for i := 1; i <= 3; i++ {
		dbErr = db.Ping()
		if dbErr != nil {
			if i < 3 {
				logger.Infof("Db connection failed, %d retry : %v", i, dbErr)
				time.Sleep(10 * time.Second)
			}
			continue
		}

		break
	}

	if dbErr != nil {
		return errors.New("can't connect to database after 3 attempts")
	}

	DbConn = db

	return nil
}
