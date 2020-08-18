// Db package for a layer for mysql package
package db

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/Antony15/gin-gonic-rest-api/mysql"
)

// Function to open a mysql database connection, or return the resource if already open
func DbConnect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/wednesday-go-test")
	DbCheckErr(err)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	err = db.PingContext(ctx)
	DbCheckErr(err)
	return db
}

/* *
 * Function to pass a query to a database & get multi row
 * */
func DbQuery(query string, args ...interface{}) (*sql.Rows, error) {
	db := DbConnect()
	defer db.Close()
	Query, err := db.Query(query, args...)
	DbCheckErr(err)
	return Query, err
}

/* *
 * Function to get number of rows of a query
 * */
func DbNumRows(query string) (int, error) {
	numrows, err := DbQuery(query)
	DbCheckErr(err)
	rowcount := 0
	for numrows.Next() {
		rowcount += 1
	}
	return rowcount, err
}

/* *
 * Function to check error
 * */
func DbCheckErr(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}
