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
    customer, err := CustomerScan(tx.QueryRow(`SELECT ID, FIRSTNAME, LASTNAME, BIRTHDATE, GENDER, EMAIL, ADDRESS FROM customer WHERE ID = ($1)`, id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Customer with id '%d' does not exist", id)
		} else {
			return nil, fmt.Errorf("Could not get customer: %v", err)
		}
	}
	return customer, nil
}

func CustomerGetForUpdate( tx * sql.Tx, id int ) (*Customer, error) {
    customer, err := CustomerScan(tx.QueryRow(`SELECT ID, FIRSTNAME, LASTNAME, BIRTHDATE, GENDER, EMAIL, ADDRESS FROM customer WHERE ID = ($1) FOR UPDATE NOWAIT`, id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Customer with id '%d' does not exist", id)
		} else {
			return nil, fmt.Errorf("Could not get customer: %v", err)
		}
	}
	return customer, nil
}

func CustomerCreate( tx * sql.Tx, customer * Customer ) (*Customer, error) {
    var id int;
	row := tx.QueryRow(`INSERT INTO customer(id, firstname, lastname, birthdate, gender, email, address) VALUES (nextval('customer_id_seq'), $1, $2, $3, $4, $5, $6) RETURNING ID`, customer.Firstname, customer.Lastname, customer.Birthdate, customer.Gender, customer.Email, customer.Address)
	if err := row.Scan(&id); err != nil {
		return nil, err
	}
	
	customer.Id = id;
    
	return customer, nil
}

func CustomerModify( tx * sql.Tx, customer * Customer ) (error) {
	_, err := tx.Exec(`UPDATE customer SET FIRSTNAME = ($2), LASTNAME = ($3), BIRTHDATE = ($4), GENDER = ($5), EMAIL = ($6), ADDRESS = ($7) WHERE ID = ($1)`, customer.Id, customer.Firstname, customer.Lastname, customer.Birthdate, customer.Gender, customer.Email, customer.Address)
	if err != nil {
		return err
	}
	
	return nil
}

func CustomerDelete( tx * sql.Tx, id int ) (error) {
	res, err := tx.Exec(`DELETE FROM customer WHERE ID = ($1)`, id)
	if err != nil {
		return err
	}
	
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return err
	}
	
	if rowCnt == 0 {
		return fmt.Errorf("Customer with id '%d' does not exist", id)
	}
	
	return nil
	
}

func CustomersFind( tx * sql.Tx, firstname string, lastname string ) ([]*Customer, error) {
    rows, err := tx.Query(`SELECT ID, FIRSTNAME, LASTNAME, BIRTHDATE, GENDER, EMAIL, ADDRESS FROM customer WHERE upper(firstname) = case when ($1) = '-' then upper(firstname) else upper($1) end AND upper(lastname) = case when ($2) = '-' then upper(lastname) else upper($2) end`, firstname, lastname )
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

func CustomersGet( tx * sql.Tx ) ([]*Customer, error) {
    rows, err := tx.Query(`SELECT ID, FIRSTNAME, LASTNAME, BIRTHDATE, GENDER, EMAIL, ADDRESS FROM customer`)
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