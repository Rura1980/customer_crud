# customer_crud
simple commandline customer crud written in golang

usage: customer_crud.exe COMMAND [ARGS]

When wrong command is supplied, usage is displayed
Only one command can be processed at once


main.go - contains command line parsing and controller functions calls
controller.go - contains supplied data checking and db calls through dto.go. DB transactions are started and completed in each controller function that needs access to db.
dto.go - contains db sqls for reading and writing customer object
db.go - contains database configurations and simple helper functions
model.go - contains Customer type definition

database.sql - contains all necessary sql for creation of DB objects. Also contains 2 sample customers.
Should be run with psql utility like this:
.\psql -a -f database.sql

error handling is very simple, no additional error types defined


Dependencies

# golang 1.10.2
# postgresql
# postgresql go library
#   go get github.com/lib/pq   