package util

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Simnek/web-go/types"
	"log"
)

func SingleRowQuery(colName, value, tableName string) (interface{}, error) {
	connStr := ConnectionString()
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database because %s", err)
	}

	var result interface{}
	var query string

	switch tableName {
	case "customer":
		var customer types.Customer
		query = fmt.Sprintf("SELECT * FROM %s WHERE %s = $1;", tableName, colName)
		row := db.QueryRowContext(context.TODO(), query, value)
		err = row.Scan(&customer.ID, &customer.Name, &customer.Address)
		result = customer
	case "order":
		var order types.Order
		query = fmt.Sprintf("SELECT * FROM %s WHERE %s = $1;", tableName, colName)
		row := db.QueryRowContext(context.TODO(), query, value)
		err = row.Scan(&order.ID, &order.ProductName, &order.Quantity, &order.Timestamp, &order.CustomerID)
		result = order
	default:
		return nil, fmt.Errorf("table name not supported")
	}

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, fmt.Errorf("unable to retrieve any data matching the query")
	case err != nil:
		return nil, fmt.Errorf("database query failed because %s", err)
	default:
		return result, nil
	}
}

func MultiRowQuery() {
	connStr := ConnectionString()
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database because %s", err)
	}

	orderQuantities := make(map[string]int)
	rows, err := db.QueryContext(context.TODO(), `SELECT food, sum(quantity) FROM "order" GROUP BY food;`)
	if err != nil {
		log.Fatalf("Database query failed because %s", err)
	}

	for rows.Next() {
		var food string
		var totalQuantity int
		err = rows.Scan(&food, &totalQuantity)
		if err != nil {
			log.Fatalf("Failed to retrieve row because %s", err)
		}
		orderQuantities[food] = totalQuantity
	}
	log.Printf("Total order quantity per food %v", orderQuantities)
}

func ParameterisedQuery(target string) {
	connStr := ConnectionString()
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database because %s", err)
	}

	var id string
	row := db.QueryRowContext(context.TODO(), `SELECT id FROM customer WHERE name = $1;`, target)
	err = row.Scan(&id)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		log.Fatalf("Unable to retrieve anyone called %s", target)
	case err != nil:
		log.Fatalf("Database query failed because %s", err)
	default:
		log.Printf("%s has an ID of %s", target, id)
	}
}

func NullTypeQuery() {
	connStr := ConnectionString()
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database because %s", err)
	}

	var allergies []sql.NullString
	rows, err := db.QueryContext(context.TODO(), `SELECT allergy FROM customer;`)
	if err != nil {
		log.Fatalf("Unable to retrieve customer allergies because %s", err)
	}

	for rows.Next() {
		var allergy sql.NullString
		err = rows.Scan(&allergy)
		if err != nil {
			log.Fatalf("Failed to scan for row because %s", err)
		}
		allergies = append(allergies, allergy)
	}
	log.Printf("Customer allergies are %v", allergies)
}
