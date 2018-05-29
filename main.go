package main

import (
    "fmt"
	"customer_crud/app"
	"os"
)

const help = `Usage: main COMMAND [ARGS]...
Commands:
	add FIRSTNAME LASTNAME BIRTHDATE<dd.mm.yyyy> GENDER<M/N> EMAIL ADDRESS - create new customer;
	del ID		   - delete customer;
	edit ID FIRSTNAME LASTNAME BIRTHDATE<dd.mm.yyyy> GENDER<M/N> EMAIL ADDRESS       - edit customer. Use '-' to preserve old value;
	show           - display all customers;
	search FIRSTNAME LASTNAME    - display records found by firstname and lastname. Use '-' to ignore one or both values. For a example "search - Putin" will search all customers with lastname Putin;
	help           - display this help.`

	
func fatal(v interface{}) {
	fmt.Println(v)
	os.Exit(1)
}

func chk(err error) {
	if err != nil {
		fatal(err)
	}
}

func show( customers []*app.Customer ) {
    for _, customer := range customers {
        showone(customer)
    }
}

func showone( customer *app.Customer ) {
    fmt.Printf("%d, %s, %s, %s, %s, %s, %s\n", customer.Id, customer.Firstname, customer.Lastname, customer.Birthdate.Format("02.01.2006"), customer.Gender, customer.Email, customer.Address)
}
 
func main() {
    if len(os.Args) < 2 {
		fatal("Usage: main COMMAND [ARG]...")
	}
	err1 := app.Connect()
	chk(err1)
	defer app.Disconnect()
	
	switch os.Args[1] {
	
	  case "add":
		if len(os.Args) != 7 && len(os.Args) != 8 {
			fatal(help)
		}
		
		var err error = nil
		firstname := os.Args[2]
		lastname := os.Args[3]
		birthdate := os.Args[4]
		gender := os.Args[5]
		email := os.Args[6]
		var address string = ""
		if len(os.Args) == 8 {
		    address = os.Args[7]
		}
	    customer, err := app.CustomerAdd( firstname, lastname, birthdate, gender, email, address)
		chk(err)
		showone(customer)	
	  case "del":
		if len(os.Args) != 3 {
			fatal(help)
		}
		err := app.CustomerRemove( os.Args[2] )
		chk(err)
	  case "edit":
		if len(os.Args) != 8 && len(os.Args) != 9 {
			fatal(help)
		}
		
		var err error = nil
		id := os.Args[2]
		firstname := os.Args[3]
		lastname := os.Args[4]
		birthdate := os.Args[5]
		gender := os.Args[6]
		email := os.Args[7]
		var address string = ""
		if len(os.Args) == 9 {
		    address = os.Args[8]
		}
		
		customer, err := app.CustomerEdit( id, firstname, lastname, birthdate, gender, email, address )
		chk(err)
		showone(customer)
	  case "show":
		if len(os.Args) != 2 {
			fatal(help)
		}
		customers, err := app.CustomerList()
		chk(err)
		show(customers)
	  case "search":
        if len(os.Args) != 4 {
			fatal(help)
		}
        customers, err := app.CustomerSearch( os.Args[2], os.Args[3] )
		chk(err)	
		show(customers)
	  case "help":
		fmt.Println(help)
	  default:
		fatal("Invalid command: " + os.Args[1])
	}
	
}