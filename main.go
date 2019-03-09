package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	if key := os.Getenv("ssh_prv_key"); key != "" {
		panic(fmt.Errorf("** ssh private key is copied into the image **"))
	}

	dbUser := "root"
	dsn := fmt.Sprintf("%v:%v@tcp(db:3306)/%v", dbUser, os.Getenv("MYSQL_ROOT_PASSWORD"), os.Getenv("MYSQL_DATABASE"))

	if db, err := sql.Open("mysql", dsn); err != nil {
		panic(err)
	} else if err = db.Ping(); err != nil {
		panic(err)
	} else {
		if _, err := db.Exec("INSERT INTO `example`(`value`) VALUES(UNIX_TIMESTAMP())"); err != nil {
			panic(err)
		}
		if res, err := db.Query("select * from `example`"); err != nil {
			panic(err)
		} else {
			fmt.Printf("example rows:\n")
			for res.Next() {
				var id int
				var value int
				if err := res.Scan(&id, &value); err != nil {
					panic(err)
				}
				fmt.Printf("-> id=%v, value=%v\n", id, value)
			}
		}
		if res, err := db.Query("SHOW TABLES"); err != nil {
			panic(err)
		} else {
			fmt.Printf("# tables:\n")
			for res.Next() {
				var tableName string
				if err := res.Scan(&tableName); err != nil {
					panic(err)
				}
				fmt.Printf("-> %v\n", tableName)
			}
		}
	}

	port := 8080
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		message := r.URL.Path
		message = strings.TrimPrefix(message, "/")
		message = "Hello " + message
		w.Write([]byte(message))
	})
	fmt.Printf("# server listening at %v\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		panic(err)
	}
}
