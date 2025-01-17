package service

import (
	"cryptoWatcher/internal/repository"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type CurrencyService struct {
	repo         *repository.CurrencyRepository
	stopChannels map[string]chan struct{}
}

func NewCurrencyService(repo *repository.CurrencyRepository) *CurrencyService {
	return &CurrencyService{repo: repo, stopChannels: make(map[string]chan struct{})}
}

func (s *CurrencyService) CurrencyAdd(coin string) error {

	price := fetchPrice(coin)

	timestamp := time.Now().Unix()

	if err := s.repo.AddCurrency(coin, price, timestamp); err != nil {
		return fmt.Errorf("failed to add currency: %w", err)
	}

	stop := make(chan struct{})
	s.stopChannels[coin] = stop

	go s.trackCurrency(coin, stop)

	log.Printf("Started tracking currency: %s", coin)
	return nil
}

func (s *CurrencyService) RemoveCurrency(coin string) error {

	if err := s.repo.RemoveCurrency(coin); err != nil {
		return fmt.Errorf("failed to remove currency: %w", err)
	}

	if stop, ok := s.stopChannels[coin]; ok {
		close(stop)
		delete(s.stopChannels, coin)
	}

	log.Println("Stopped tracking currency: %s", coin)
	return nil
}

func (s *CurrencyService) GetPrice(coin string, timestamp int64) (float64, error) {
	exists, err := s.repo.IsCurrencyTracked(coin)
	if err != nil {
		return 0, fmt.Errorf("failed to check if currency is tracked: %w", err)
	}
	if !exists {
		return 0, fmt.Errorf("currency %s does not exists", coin)
	}

	price, err := s.repo.GetPrice(coin, timestamp)
	if err != nil {
		return 0, fmt.Errorf("failed to get price: %w", err)
	}

	return price, nil
}

func (s *CurrencyService) trackCurrency(coin string, stop chan struct{}) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			price := fetchPrice(coin)
			timestamp := time.Now().Unix()

			if err := s.repo.AddCurrency(coin, price, timestamp); err != nil {
				log.Printf("Failed to update price for %s: %v", coin, err)
			} else {
				log.Printf("Updated price for %s: %f", coin, price)
			}
		case <-stop:
			log.Printf("Stopped tracking currency: %s", coin)
			return
		}
	}
}

func fetchPrice(coin string) float64 {
	apiKey := "CG-TQdoUcTNZYBWpTNtZNzfv7Et"
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=usd&api_key=%s", coin, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return -1
	}
	defer resp.Body.Close()

	var result map[string]map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Printf("Failed to decode response: %v", err)
		return -1
	}

	priceData, ok := result[coin]["usd"]
	if !ok {
		log.Printf("Price not found for %s", coin)

	}

	var price float64
	switch v := priceData.(type) {
	case float64:
		price = v
	case string:
		price, err = strconv.ParseFloat(v, 64)
		if err != nil {
			log.Printf("Failed to parse price for %s: %v", coin, err)
			return -1
		}
	default:
		log.Printf("Unexpected price format for %s: %T", coin, v)
		return -1
	}

	return price
}
