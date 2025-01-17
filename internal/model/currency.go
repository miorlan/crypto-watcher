package model

// Currency представляет информацию о валюте.
// @Description Информация о валюте, включая её название, цену и временную метку.
type Currency struct {
	// Уникальный идентификатор валюты.
	ID int `json:"id"`

	// Название валюты.
	Coin string `json:"coin"`

	// Текущая цена валюты.
	Price float64 `json:"price"`

	// Временная метка (timestamp) записи.
	Timestamp int64 `json:"timestamp"`
}

// CurrencyRequest представляет информацию о валюте.
type CurrencyRequest struct {

	// Название валюты.
	Coin string `json:"coin"`
}

// GetPriceRequest запрос на получение стоимости валюты по времени.
type GetPriceRequest struct {

	// Название валюты.
	Coin string `json:"coin"`

	// Временная метка (timestamp) записи.
	Timestamp int64 `json:"timestamp"`
}
