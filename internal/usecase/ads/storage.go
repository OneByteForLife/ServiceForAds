package ads

import (
	"ServiceForAds/internal/entity"
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type AdsStorage struct {
	pool *pgxpool.Pool
}

func NewStorage(pool *pgxpool.Pool) Storage {
	return &AdsStorage{pool}
}

const (
	advertisements = "advertisements" // Таблица с товарами
)

// Выборка по id объявленияw
func (s *AdsStorage) GetOne(id int) (entity.Advertisements, error) {
	var ads entity.Advertisements

	query, _, _ := goqu.From(advertisements).Where(goqu.ExOr{
		"id": id,
	}).ToSQL()

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
		query, _, _ = goqu.From(advertisements).Limit(uint(limit)).Offset(uint(offset)).ToSQL()
	}

	switch {
	case sortType == "asc":
		query, _, _ = goqu.From(advertisements).OrderAppend(goqu.I(sortBy).Asc()).Limit(uint(limit)).Offset(uint(offset)).ToSQL()
	case sortType == "desc":
		query, _, _ = goqu.From(advertisements).OrderAppend(goqu.I(sortBy).Asc()).Limit(uint(limit)).Offset(uint(offset)).ToSQL()
	}

	rows, err := s.pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

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
	ds := goqu.Insert(advertisements).
		Cols("product_name", "product_description", "product_main_picture", "product_more_pictures", "date_create", "price").
		Vals(
			goqu.Vals{
				&ads.ProductName,
				&ads.Description,
				&ads.MainPicture,
				pq.Array(&ads.MorePictures), // Пришлось юзать pq.Array из за того что goqu.Insert не правильно преобразовывал массивы
				&ads.DateCreate, &ads.Price,
			})

	query, _, _ := ds.ToSQL()

	_, err := s.pool.Exec(context.Background(), query)
	if err != nil {
		return err
	}

	return nil
}
