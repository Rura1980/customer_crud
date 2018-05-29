package app

import (
	"time"
	"regexp"
	"strconv"
)

func CustomerCheck( firstname string, lastname string, birthdate string, gender string, email string, address string ) (*Customer, error) {
    if firstname == "" || len(firstname) > 100 {
	    return nil, BusinessError{"Customer firstname is not suitable"}
	}
	
	if lastname == "" || len(lastname) > 100 {
	    return nil, BusinessError{"Customer lastname is not suitable"}
	}
	
	d, err := time.Parse( "02.01.2006", birthdate)
	
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

func CustomerEdit( id string, firstname string, lastname string, birthdate string, gender string, email string, address string ) ( *Customer, error) {
	id_, err := strconv.Atoi(id)
	if err != nil {
	    return nil, err
	}
	
	tx, err1 := appDb.Begin()
	if err1 != nil {
	    return nil, err1
	}
	defer tx.Rollback()
	
	customer, err2 := CustomerGetForUpdate(tx, id_)
	if err2 != nil {
	    return nil, err2
	}
	
	var firstname_ string
	if firstname != "-" {
	    firstname_ = firstname
	} else {
	    firstname_ = customer.Firstname
	}
	
	var lastname_ string
	if lastname != "-" {
	    lastname_ = lastname
	} else {
	    lastname_ = customer.Lastname
	}
	
	var birthdate_ string
	if birthdate != "-" {
	    birthdate_ = birthdate
	} else {
	    birthdate_ = customer.Birthdate.Format("02.01.2006")
	}
	
	var gender_ string
	if gender != "-" {
	    gender_ = gender
	} else {
		if customer.Gender == "Male" {
		    gender_ = "M"
		} else {
		    gender_ = "F"
		}
	}
	
	var email_ string
	if email != "-" {
	    email_ = email
	} else {
	    email_ = customer.Email
	}
	
	var address_ string
	if address != "-" {
	    address_ = address
	} else {
	    address_ = customer.Address
	}
	
	customer, err3 := CustomerCheck( firstname_, lastname_, birthdate_, gender_, email_, address_ )
	if err3 != nil {
	    return nil, err3
	}
	
	customer.Id = id_
	
	err4 := CustomerModify(tx, customer)
	if err4 != nil {
	    return nil, err4
	}
	tx.Commit()
	return customer, nil
}

func CustomerRemove( id string ) (error) {
	id_, err := strconv.Atoi(id)
	if err != nil {
	    return err
	}
	
    tx, err1 := appDb.Begin()
	if err1 != nil {
	    return err1
	}
	defer tx.Rollback()
	
	err2 := CustomerDelete( tx, id_ )
	if err2 != nil {
	    return err2
	}
	tx.Commit()
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

func CustomerSearch( firstname string, lastname string ) ([]*Customer, error) {
    tx, err := appDb.Begin()
	if err != nil {
	    return nil, err
	}
	defer tx.Rollback()
	
	customers, err := CustomersFind( tx, firstname, lastname )
	tx.Commit()
    return customers, nil	
}