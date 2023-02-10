package database

import (
	"ServiceForAds/config"
	"context"
	"fmt"
	"time"

	pgxpool "github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

func ConnectToDatabase(conf *config.Config) *pgxpool.Pool {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s&pool_max_conns=50", conf.Database.Username, conf.Database.Password, conf.Database.Host, conf.Database.Port, conf.Database.Dbname, conf.Database.Sslmode)
	pool, err := pgxpool.Connect(context.Background(), url)
	if err != nil {
		logrus.Errorf("database connection failed: %v [starting reconnect]", err)
		time.Sleep(time.Second * 10)
		ConnectToDatabase(conf)
	} else {
		logrus.Info("Success connect to database")
	}
	return pool
}
