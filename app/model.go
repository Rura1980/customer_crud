package app

import (
	"time"
)

type Customer struct {
    Id 		  int
	Firstname string
	Lastname  string
	Birthdate time.Time
    Gender 	  string
    Email     string
    Address   string
}