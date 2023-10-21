package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/minguu42/myapp"
)

func main() {
	db := myapp.OpenDB(os.Getenv("DRIVER"), os.Getenv("DSN"))
	defer myapp.CloseDB(db)

	if err := db.Ping(); err != nil {
		log.Fatal("db.Ping failed:", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, "Hello, 世界！")
	})

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
