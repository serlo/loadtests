package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var pageIds = []int{18659, 18676, 18695, 18719, 18732, 18739, 18747, 18764, 18766, 18773, 18775, 18815, 18822, 18824, 18826, 18835, 18839, 18844, 18861, 18872, 18939, 18941, 18956, 18959, 18987, 18988, 18995, 19001, 19006, 19013, 19016, 19032, 19037, 19044, 19046, 19059, 19071, 19079, 19080, 19098, 19104, 19124, 19125, 19131, 19143, 19148, 19150, 19179, 19190, 19211, 19212, 19214, 19229, 19252, 19270, 19290, 19294, 19297, 19328, 19332, 19334, 19352, 19377, 19380, 19481, 19485, 19487, 19490, 19492, 19495, 19502, 19515, 19519, 19535, 19537, 19539, 19543, 19549, 19554}
var username string
var password string
var port = 3306

func main() {
	portVar := os.Getenv("TEST_DB_PORT")
	if portVar != "" {
		var err error
		port, err = strconv.Atoi(portVar)
		if err != nil {
			fmt.Printf("TEST_DB_PORT not an integer")
			os.Exit(1)
		}
		fmt.Printf("use db port [%d]\n", port)
	}
	username = os.Getenv("TEST_DB_USERNAME")
	if username == "" {
		fmt.Printf("TEST_DB_USERNAME variable not set")
		os.Exit(1)
	}
	password = os.Getenv("TEST_DB_PASSWORD")
	if password == "" {
		fmt.Printf("TEST_DB_PASSWORD variable not set")
		os.Exit(1)
	}

	testNoPool()
	testWithPool()
}

func testNoPool() {
	totalMs := time.Duration(0)
	for _, id := range pageIds {
		startTime := time.Now()
		db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(127.0.0.1:%d)/serlo", username, password, port))
		// if there is an error opening the connection, handle it
		if err != nil {
			fmt.Printf("ERROR [%s]", err.Error())
			os.Exit(1)
		}

		stmt, _ := db.Prepare("select id, type_id from entity where id = ?")
		query(stmt, id)
		db.Close()
		responseTime := time.Since(startTime) / time.Millisecond
		totalMs += responseTime
	}
	fmt.Printf("{ \"avg\": %d }\n", totalMs/time.Duration(len(pageIds)))
}

func testWithPool() {
	totalMs := time.Duration(0)

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(127.0.0.1:%d)/serlo", username, password, port))
	// if there is an error opening the connection, handle it
	if err != nil {
		fmt.Printf("ERROR [%s]", err.Error())
		os.Exit(1)
	}
	defer db.Close()

	stmt, _ := db.Prepare("select id, type_id from entity where id = ?")
	for _, id := range pageIds {
		startTime := time.Now()
		query(stmt, id)
		responseTime := time.Since(startTime) / time.Millisecond
		totalMs += responseTime
	}
	fmt.Printf("{ \"avg\": %d }\n", totalMs/time.Duration(len(pageIds)))
}

func query(stmt *sql.Stmt, id int) {
	rows, err := stmt.Query(id)
	if err != nil {
		fmt.Printf("ERROR [%s]", err.Error())
		os.Exit(1)
	}

	defer rows.Close()

	typeID := 0
	for rows.Next() {
		id := 0
		err = rows.Scan(&id, &typeID)
	}
}
