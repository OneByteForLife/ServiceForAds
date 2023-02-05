package database

import (
	"ServiceForAds/config"
	"context"
	"fmt"

	pgxpool "github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

func ConnectToDatabase(conf *config.Config) (*pgxpool.Pool, error) {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s&pool_max_conns=50", conf.Database.Username, conf.Database.Password, conf.Database.Host, conf.Database.Port, conf.Database.Dbname, conf.Database.Sslmode)

	pool, err := pgxpool.Connect(context.Background(), url)
	if err != nil {
		logrus.Errorf("Err connect to database - %s", err)
		return nil, err
	}

	return pool, nil
}
