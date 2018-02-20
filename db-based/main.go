package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/karampok/go-cf-app/backend/db"
)

var (
	sqldb *sql.DB
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	sqldb, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Can not open DB:", err)
	}
	defer sqldb.Close()

	var version string
	sqldb.QueryRow("SELECT VERSION()").Scan(&version)

	if err := db.DeployV1Schema(sqldb); err != nil {
		log.Printf("Can not connect to db\n")
		log.Fatal("Can not apply schema :", err)
	}

	log.Printf("Connected to: %s\n", version)
	log.Printf("Schema v1 OK")

	systemPort, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("invalid required env var PORT")
	}

	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", systemPort), mux)
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%v: %v\n", item, price)
	}
}

type database map[string]int

func (db database) price(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case "GET":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%v\n", price)
		return

	case "POST":
		item := req.URL.Query().Get("item")
		priceString := req.URL.Query().Get("price")
		if price, err := strconv.Atoi(priceString); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "not valid price: %q\n", priceString)
			return
		} else {
			if _, err := sqldb.Exec(`INSERT INTO products (name, value) VALUES ("socks",5)`); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "error: %d%vhn", item, price)
				return
			}
			fmt.Fprintf(w, "new price %s\n", priceString)
		}
		return

	case "DELETE":
		return
	}

}
