package app

import (
    "database/sql"
	"time"
	"fmt"
	_ "github.com/lib/pq"
)

func CustomerScan(s rowScanner) (*Customer, error) {
	var (
		id int
		firstname string
		lastname string
		birthdate time.Time
		gender string
		email string
		address sql.NullString
	)
	if err := s.Scan(&id, &firstname, &lastname, &birthdate, &gender, &email, &address); err != nil {
		return nil, err
	}

	customer := &Customer{ id, firstname, lastname, birthdate, gender, email, NullString2String(address) }
	return customer, nil
}

func CustomerGet( tx * sql.Tx, id int ) (*Customer, error) {
    customer, err := CustomerScan(tx.QueryRow(`SELECT * FROM customer WHERE id = ($1)`, id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Customer with id '%d' does not exist", id)
		} else {
			return nil, fmt.Errorf("Could not get player: %v", err)
		}
	}
	return customer, nil
}

func CustomerCreate( tx * sql.Tx, customer * Customer ) (*Customer, error) {
    var id int;
	row := tx.QueryRow(`INSERT INTO customer(id, firstname, lastname, birthdate, gender, email, address) VALUES (nextval('customer_id_seq'), $1, $2, $3, $4, $5, $6) RETURNING id`, customer.Firstname, customer.Lastname, customer.Birthdate, customer.Gender, customer.Email, customer.Address)
	if err := row.Scan(&id); err != nil {
		return nil, err
	}
	
	customer.Id = id;
    
	return customer, nil
}

func CustomersGet( tx * sql.Tx ) ([]*Customer, error) {
    rows, err := tx.Query(`SELECT * FROM customer`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var customers []*Customer
	
	for rows.Next() {
		customer, err := CustomerScan(rows)
		if err != nil {
			return nil, err
		}

		customers = append(customers, customer)
	}

	return customers, nil
}