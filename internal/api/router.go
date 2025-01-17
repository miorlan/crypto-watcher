package api

import (
	"cryptoWatcher/internal/api/handler"
	"cryptoWatcher/internal/database"
	"cryptoWatcher/internal/repository"
	"cryptoWatcher/internal/service"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "cryptoWatcher/docs"
)

func SetupRouter(db *database.DB) *chi.Mux {
	currencyRepo := repository.NewCurrencyRepository(db.DB)

	currencyService := service.NewCurrencyService(currencyRepo)

	currencyHandler := handler.NewCurrencyHandler(currencyService)

	r := chi.NewRouter()

	r.Get("/swagger/*", httpSwagger.WrapHandler)
	r.Post("/currency/add", currencyHandler.AddCurrency)
	r.Post("/currency/remove", currencyHandler.RemoveCurrency)
	r.Get("/currency/price", currencyHandler.GetPrice)

	return r
}
