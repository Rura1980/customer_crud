package app

import (
	"fmt"
	"time"
	"regexp"
)

func CustomerCheck( firstname string, lastname string, birthdate string, gender string, email string, address string ) (*Customer, error) {
    if firstname == "" || len(firstname) > 100 {
	    return nil, BusinessError{"Customer firstname is not suitable"}
	}
	
	if lastname == "" || len(lastname) > 100 {
	    return nil, BusinessError{"Customer lastname is not suitable"}
	}
	
	fmt.Println(birthdate)
	const shortForm = "02.01.2006"
	d, err := time.Parse( shortForm, birthdate)
	
	if err != nil {
	    return nil, err
	}
	   
	age := time.Now().Year() - d.Year()
	
	if age < 18 || age > 60 {
	    return nil, BusinessError{"Customer is too young or too old"}
	}  
	
	if gender != "M" && gender != "F" {
	   return nil, BusinessError{"Customer has unknown gender"}
	}
	
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !re.MatchString(email) {
	    return nil, BusinessError{"Customer email format is wrong"}
	}
	
	if len(address) > 200 {
	    return nil, BusinessError{"Customer address is too long"}
	}
	
	var gender_ string
	
	if gender == "M" {
	    gender_ = "Male"
	}
	if gender == "F" {
        gender_ = "Female"	
	}
	return &Customer{-1, firstname, lastname, d, gender_, email, address}, nil
} 

func CustomerAdd( firstname string, lastname string, birthdate string, gender string, email string, address string ) (*Customer, error) {
    customer, err := CustomerCheck( firstname, lastname, birthdate, gender, email, address )
	if err != nil {
	    return nil, err
	}
	
	tx, err1 := appDb.Begin()
	if err1 != nil {
	    return nil, err1
	}
	defer tx.Rollback()
	
	customer, err2 := CustomerCreate(tx, customer)
	if err2 != nil {
	    return nil, err2
	}
	tx.Commit()
	return customer, nil
}

func CustomerEdit( ) (error) {
    return nil
}

func CustomerDelete( ) (error) {
    return nil
}

func CustomerList( ) ([]*Customer, error) {
    tx, err := appDb.Begin()
	if err != nil {
	    return nil, err
	}
	defer tx.Rollback()
	
	customers, err := CustomersGet( tx )
	if err != nil {
	    return nil, err
	}
	tx.Commit()
    return customers, nil
}

func CustomerSearch( ) (error) {
    return nil
}