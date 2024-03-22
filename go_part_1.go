package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "net/http"
    "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func initDB() {
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    db, err = sql.Open("mysql", dbUser+":"+dbPassword+"@/"+dbName)
    if err != nil {
        log.Fatal(err)
    }


err = db.Ping()
if err != nil {
    log.Fatal(err)
    }
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }

func handleSearch(w http.ResponseWriter, r *http.Request) {
    searchQuery := r.URL.Query().Get("query")
    if searchQuery == "" {
        http.Error(w, "Query parameter is missing", http.StatusBadRequest)
        return
    }

query := "SELECT * FROM products WHERE name LIKE '%' || $1 || '%'"
rows, err := db.Query(query, searchQuery)

if err != nil {
    http.Error(w, "Query failed", http.StatusInternalServerError)
    log.Println(err)
    return
}
defer rows.Close()

var products []string
for rows.Next() {
    var name string
    err := rows.Scan(&name)
    if err != nil {
        log.Fatal(err)
    }
    products = append(products, name)
}

fmt.Fprintf(w, "Found products: %v\n", products)
}

func main() {
    initDB()
    defer db.Close()

http.HandleFunc("/search", searchHandler)
fmt.Println("Server is running")
log.Fatal(http.ListenAndServe(":8080", nil))
}