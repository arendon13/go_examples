package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Store organizes store data
type Store struct {
	StoreID   int
	StoreNum  int
	StoreName string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const createStores = "CREATE TABLE tbl_Stores( " +
	"	StoreID     int NOT NULL auto_increment, " +
	"	StoreNum    int, " +
	"	StoreName   varchar(255), " +
	"	PRIMARY KEY (StoreID) " +
	")"

const storeData = "INSERT INTO tbl_Stores(StoreNum, StoreName) VALUES " +
	"(1,'STORE #1'), (2,'STORE #2'), (3,'STORE #3'), (4,'STORE #4'), (5,'STORE #5')"

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(172.17.0.1:3306)/lpts_test")
	check(err)
	defer db.Close()

	stmt, err := db.Prepare(createStores)
	check(err)
	defer stmt.Close()

	_, err = stmt.Exec()
	check(err)

	// stmt, err = db.Prepare(storeData)
	// check(err)
	// defer stmt.Close()

	// _, err = stmt.Exec()
	// check(err)

	// // insert row
	// stmtIn, err := db.Prepare("INSERT INTO tbl_Stores(StoreNum, StoreName) VALUES( ?, ? )")
	// check(err)
	// defer stmtIn.Close()
	// _, err = stmtIn.Exec(69, "JONS #69")
	// check(err)

	// // select multiple rows
	// maxID := 0
	// rows, err := db.Query("SELECT * FROM tbl_Stores")
	// check(err)
	// defer rows.Close()
	// for rows.Next() {
	// 	store := Store{}
	// 	err := rows.Scan(&store.StoreID, &store.StoreNum, &store.StoreName)
	// 	check(err)
	// 	maxID = store.StoreID
	// 	fmt.Println(store)
	// }

	// // select single row
	// id := 4
	// var storeName string
	// stmtRead, err := db.Prepare("SELECT StoreName FROM tbl_Stores WHERE StoreID = ?")
	// check(err)
	// defer stmtRead.Close()
	// err = stmtRead.QueryRow(id).Scan(&storeName)
	// check(err)
	// fmt.Printf("\n%s\n", storeName)

	// // delete row
	// stmtDel, err := db.Prepare("DELETE FROM tbl_Stores WHERE StoreID = ?")
	// check(err)
	// defer stmtDel.Close()
	// _, err = stmtDel.Exec(maxID)
	// check(err)

	// // select multiple rows to show delete effects
	// rows, err = db.Query("SELECT * FROM tbl_Stores")
	// check(err)
	// defer rows.Close()
	// fmt.Println()
	// for rows.Next() {
	// 	store := Store{}
	// 	err := rows.Scan(&store.StoreID, &store.StoreNum, &store.StoreName)
	// 	check(err)
	// 	fmt.Println(store)
	// }
}
