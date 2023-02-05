package ads

import (
	"ServiceForAds/internal/entity"
	"ServiceForAds/tools/conv"
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
	ads, err := s.storage.GetOne(conv.ConvertQuery(id))
	if err != nil {
		logrus.Errorf("error in receiving the ad: %v", err)
		return ads, fmt.Errorf("sorry but the ad with the %s was not found", id)
	}

	return ads, nil
}

func (s *AdsService) GetAll(limit string, offset string) ([]entity.Advertisements, error) {
	ads, err := s.storage.GetAll(conv.ConvertQuery(limit), conv.ConvertQuery(offset))
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
		return errors.New("ad creation error")
	}
	return nil
}

// Для сериализации json'а в структуру
func unmarsalObject(data []byte) (entity.Advertisements, error) {
	var ads entity.Advertisements

	if err := json.Unmarshal(data, &ads); err != nil {
		logrus.Errorf("error wrong data format ad: %s", err)
		return ads, errors.New("err wrong data format for ads")
	}

	// Валидация полей

	t, err := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		logrus.Errorf("error parce date from ads")
		return ads, nil
	}

	ads.DateCreate = t
	return ads, nil
}
