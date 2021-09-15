package main_test

import (
	"log"
	"os"
	"testing"

	main "github.com/Mfdsix/go-postgre"
)

var a main.App

func TestMain(m *testing.M) {
	a.Initialize(
		main.GetEnvVal("APP_DB_USERNAME"),
		main.GetEnvVal("APP_DB_PASSWORD"),
		main.GetEnvVal("APP_DB_NAME"),
	)

	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreatinoQuery); err != nil {
		log.Fatalf("failed execute table creation, err: %v", err)
		log.Fatal()
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM products")
	a.DB.Exec("ALTER SEQUENCE product_id_seq RESTART WITH 1")
}

const tableCreatinoQuery = `CREATE TABLE IF NOT EXISTS products
(
	id SERIAL,
	name TEXT NOT NULL,
	price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
	CONSTRAINT products_pkey PRIMARY KEY (id)
)`
