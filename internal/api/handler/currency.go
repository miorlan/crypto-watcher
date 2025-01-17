// @title CryptoWatcher API
// @version 1.0
// @description API для отслеживания стоимости криптовалют.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
package handler

import (
	"cryptoWatcher/internal/model"
	"cryptoWatcher/internal/service"
	"encoding/json"
	"net/http"
)

type CurrencyHandler struct {
	currencyService *service.CurrencyService
}

func NewCurrencyHandler(currencyService *service.CurrencyService) *CurrencyHandler {
	return &CurrencyHandler{currencyService: currencyService}
}

// AddCurrency godoc
// @Summary Добавить валюту
// @Description Добавляет новую валюту для отслеживания.
// @Tags currency
// @Accept json
// @Produce json
// @Param request body model.CurrencyRequest true "Данные для добавления валюты" example value({"coin": "dogecoin"})
// @Success 200 {object} map[string]string "Успешный ответ" example({"message": "Currency added"})
// @Failure 400 {string} string "Неверный запрос"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /currency/add [post]
func (h *CurrencyHandler) AddCurrency(w http.ResponseWriter, r *http.Request) {
	var request model.CurrencyRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if err := h.currencyService.CurrencyAdd(request.Coin); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Currency added"})
}

// RemoveCurrency godoc
// @Summary Удалить валюту
// @Description Удаляет валюту из списка отслеживаемых.
// @Tags currency
// @Accept json
// @Produce json
// @Param request body model.CurrencyRequest true "Данные для удаления валюты" example({"coin": "dogecoin"})
// @Success 200 {object} map[string]string "Успешный ответ" example({"message": "Currency removed"})
// @Failure 400 {string} string "Неверный запрос"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /currency/remove [post]
func (h *CurrencyHandler) RemoveCurrency(w http.ResponseWriter, r *http.Request) {
	var request model.CurrencyRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if err := h.currencyService.RemoveCurrency(request.Coin); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Currency removed"})
}

// GetPrice godoc
// @Summary Получить цену валюты
// @Description Возвращает цену валюты на указанный timestamp.
// @Tags currency
// @Accept json
// @Produce json
// @Param request body model.GetPriceRequest true "Данные для получения цены" example({"coin": "bitcoin", "timestamp": 1737039590})
// @Success 200 {object} map[string]float64 "Успешный ответ" example({"price": 97675})
// @Failure 400 {string} string "Неверный запрос"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /currency/price [get]
func (h *CurrencyHandler) GetPrice(w http.ResponseWriter, r *http.Request) {
	var request model.GetPriceRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	price, err := h.currencyService.GetPrice(request.Coin, request.Timestamp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]float64{"price": price})

}
