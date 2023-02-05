package entity

import "time"

/*
	Название (Обязательно)
	Описание (Обязательно размер = 100 - 1000)
	Ссылка на фото основное (Обязательно)
	Ссылки на фото дополнительные (Не обязательно)
	Цена (Обязательно)
*/

// Описание модели объявлений
type Advertisements struct {
	ID           int       `json:"id,omitempty"`
	ProductName  string    `json:"product_name"`
	Description  string    `json:"description"`
	MainPicture  string    `json:"main_picture"`
	MorePictures []string  `json:"second_pictures,omitempty"`
	DateCreate   time.Time `json:"date_create,omitempty"`
	Price        float64   `json:"product_price"`
}
