package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	db, err := sql.Open("mysql", "root:@/lpts")
	check(err)
	defer db.Close()

	stmtOut, err := db.Prepare("SELECT * FROM tbl_Stores")
	check(err)

	// stmtOut.QueryRow
	defer stmtOut.Close()

}
