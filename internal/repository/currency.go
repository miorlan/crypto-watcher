package repository

import (
	"database/sql"
	"fmt"
)

type CurrencyRepository struct {
	db *sql.DB
}

func NewCurrencyRepository(db *sql.DB) *CurrencyRepository {
	return &CurrencyRepository{db: db}
}

func (r *CurrencyRepository) AddCurrency(coin string, price float64, timestamp int64) error {
	query := "INSERT INTO prices (coin, price, timestamp) VALUES ($1, $2, $3)"

	if price < 0 {
		return nil
	}
	_, err := r.db.Exec(query, coin, price, timestamp)
	if err != nil {
		return fmt.Errorf("failed to add currency: %w", err)
	}
	return nil
}

func (r *CurrencyRepository) RemoveCurrency(coin string) error {
	query := "DELETE FROM prices WHERE coin=$1"
	_, err := r.db.Exec(query, coin)
	if err != nil {
		return fmt.Errorf("failed to remove currency: %w", err)
	}
	return nil
}

func (r *CurrencyRepository) GetPrice(coin string, timestamp int64) (float64, error) {

	var price float64

	query := "SELECT price FROM prices WHERE coin = $1 AND timestamp <= $2"

	err := r.db.QueryRow(query, coin, timestamp).Scan(&price)
	if err != nil {
		return 0.0, fmt.Errorf("failed to get price: %w", err)
	}

	return price, nil
}

func (r *CurrencyRepository) IsCurrencyTracked(coin string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM prices WHERE coin=$1)"

	err := r.db.QueryRow(query, coin).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check if currency is tracked: %w", err)
	}
	return exists, nil
}
