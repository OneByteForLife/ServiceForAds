package validation

import (
	"errors"
	"fmt"
	"strings"
)

// Для валидации url фотографий
func ValideteUrlPicture(url string) error {
	if !strings.Contains(url, "http://") || !strings.Contains(url, "https://") {
		return errors.New("url addr is not valid")
	}

	if len(url) <= 0 || len(url) >= 1000 {
		return errors.New("the length of the address exceeds the allowable length")
	}

	var shareds []string = []string{".jpg", ".jpeg", ".png"}
	for idx, shared := range shareds {
		if !strings.Contains(url, shared) {
			return fmt.Errorf("the %s of the photo under index %d is not valid", shared, idx)
		}
	}

	return nil
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
