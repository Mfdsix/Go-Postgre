package main

import (
	"database/sql"
	"errors"
)

type MProduct struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *MProduct) getProduct(db *sql.DB) error {
	return errors.New("Not Implemented")
}

func (p *MProduct) updateProduct(db *sql.DB) error {
	return errors.New("Not Implemented")
}

func (p *MProduct) deleteProduct(db *sql.DB) error {
	return errors.New("Not Implemented")
}

func (p *MProduct) createProduct(db *sql.DB) error {
	return errors.New("Not Implemented")
}

func getProducts(db *sql.DB, start int, count int) ([]MProduct, error) {
	return nil, errors.New("Not Implemented")
}
