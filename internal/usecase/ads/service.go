package ads

import (
	"ServiceForAds/internal/entity"
	"ServiceForAds/utils/conv"
	"ServiceForAds/utils/validation"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type AdsService struct {
	storage Storage
}

func NewService(s Storage) Service {
	return &AdsService{s}
}

func (s *AdsService) GetOne(id string) (entity.Advertisements, error) {
	if id == "" {
		return entity.Advertisements{}, errors.New("error id cannot be empty")
	}

	ads, err := s.storage.GetOne(conv.ConvertQuery(id))
	if err != nil {
		logrus.Errorf("error in receiving the ad: %v", err)
		return ads, fmt.Errorf("sorry but the ad with the %s was not found", id)
	}

	return ads, nil
}

func (s *AdsService) GetAll(limit string, offset string, sortBy string, sortType string) ([]entity.Advertisements, error) {

	if err := validation.ValidateUrlQuery(limit, offset, sortBy, sortType); err != nil {
		return nil, err
	}

	ads, err := s.storage.GetAll(conv.ConvertQuery(limit), conv.ConvertQuery(offset), sortBy, sortType)
	if err != nil {
		logrus.Errorf("error in ad issuance: %v", err)
		return nil, errors.New("error in ad issuance check the parameters of the request")
	}

	if ads == nil {
		return ads, errors.New("out of range, check query parameters")
	}

	return ads, nil
}

func (s *AdsService) Create(body []byte) error {
	ads, err := unmarsalObject(body)
	if err != nil {
		return err
	}

	err = s.storage.Create(ads)
	if err != nil {
		logrus.Errorf("ad creation error: %v", err)
		return errors.New("ad creation error check your json obj")
	}
	return nil
}

func unmarsalObject(data []byte) (entity.Advertisements, error) {
	var ads entity.Advertisements

	if err := json.Unmarshal(data, &ads); err != nil {
		logrus.Errorf("error parce data from ads: %s", err)
		return ads, errors.New("err wrong data format for ads")
	}

	// Валидация основного изображения
	if err := validation.ValideteUrlPicture(ads.MainPicture); err != nil {
		return ads, fmt.Errorf("error the link to the main image is not correct: %v", err)
	}

	// Валидация ссылок на изображения
	for idx, val := range ads.MorePictures {
		if err := validation.ValideteUrlPicture(val); err != nil {
			return ads, fmt.Errorf("error link number %d is wrong: %v", idx+1, err)
		}
	}

	// Валидация описания товара
	if err := validation.ValideteText(ads.Description); err != nil {
		return ads, fmt.Errorf("error description of your product does not fit into the acceptable limits (200 - 1000): %v", err)
	}

	// Формирование времени для объявления
	t, err := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		logrus.Errorf("error parce date from ads: %v", err)
		return ads, nil
	}
	ads.DateCreate = t

	return ads, nil
}
