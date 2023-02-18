package validation

import (
	"errors"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	date  = "date_create"
	price = "price"
	asc   = "asc"
	desc  = "desc"
)

// Для валидации url фотографий
func ValideteUrlPicture(url string) error {
	if !strings.Contains(url, "http://") && !strings.Contains(url, "https://") {
		return fmt.Errorf("%s is wrong", url)
	}

	if len(url) <= 0 && len(url) > 1000 {
		return errors.New("the length of the address exceeds the allowable length")
	}

	switch {
	case filepath.Ext(url) == ".jpg":
		return nil
	case filepath.Ext(url) == ".jpeg":
		return nil
	case filepath.Ext(url) == ".png":
		return nil
	}
	return fmt.Errorf("error the link to the %s is not correct", url)
}

// Для валидации названия и описания
func ValideteText(text string) error {
	if len(text) < 200 {
		return errors.New("product description must have more than 200 characters")
	}

	if len(text) > 1000 {
		return errors.New("product description must have more than 1000 characters")
	}

	return nil
}

// Валидация url параметров
func ValidateUrlQuery(limit string, offset string, sortBy string, sortType string) error {
	if limit == "" && offset == "" {
		return errors.New("error parameters for pagination are mandatory")
	}

	if sortBy == "" && sortType == "" {
		return nil
	}

	if _, err := strconv.Atoi(limit); err != nil {
		return err
	}

	if _, err := strconv.Atoi(offset); err != nil {
		return err
	}

	if sortBy != date && sortBy != price {
		return fmt.Errorf("error params for sorting - %s", sortBy)
	}

	if sortType != asc && sortType != desc {
		return fmt.Errorf("error type for sorting - %s", sortType)
	}

	return nil
}
