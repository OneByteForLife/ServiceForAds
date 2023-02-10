package ads

import (
	"ServiceForAds/internal/entity"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type AdsStorage struct {
	pool *pgxpool.Pool
}

func NewStorage(pool *pgxpool.Pool) Storage {
	return &AdsStorage{pool}
}

// Выборка по id объявленияw
func (s *AdsStorage) GetOne(id int) (entity.Advertisements, error) {
	var ads entity.Advertisements

	query := fmt.Sprintf("SELECT * FROM advertisements WHERE id = %d", id)
	err := s.pool.QueryRow(context.Background(), query).Scan(&ads.ID, &ads.ProductName, &ads.Description, &ads.MainPicture, &ads.MorePictures, &ads.DateCreate, &ads.Price)
	if err != nil {
		return ads, err
	}

	return ads, nil
}

// Выборка всех объявлений
func (s *AdsStorage) GetAll(limit int, offset int, sortBy string, sortType string) ([]entity.Advertisements, error) {
	var ads []entity.Advertisements

	var query string
	if sortBy == "" && sortType == "" {
		query = fmt.Sprintf("SELECT * FROM advertisements LIMIT %d OFFSET %d", limit, offset)
	} else {
		query = fmt.Sprintf("SELECT * FROM advertisements ORDER BY %s %s LIMIT %d OFFSET %d ", sortBy, sortType, limit, offset)
	}

	rows, err := s.pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	var temp entity.Advertisements
	for rows.Next() {
		err := rows.Scan(&temp.ID, &temp.ProductName, &temp.Description, &temp.MainPicture, &temp.MorePictures, &temp.DateCreate, &temp.Price)
		if err != nil {
			logrus.Errorf("err scan row to temp: %s", err)
			return nil, err
		}
		ads = append(ads, temp)
	}
	return ads, nil
}

// Добавления объявления
func (s *AdsStorage) Create(ads entity.Advertisements) error {
	_, err := s.pool.Exec(context.Background(), "INSERT INTO advertisements (product_name, product_description, product_main_picture, product_more_pictures, date_create, price) VALUES ($1, $2, $3, $4, $5, $6)",
		ads.ProductName,
		ads.Description,
		ads.MainPicture,
		ads.MorePictures,
		ads.DateCreate,
		ads.Price,
	)
	if err != nil {
		return err
	}
	return nil
}
