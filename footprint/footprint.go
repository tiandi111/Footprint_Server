package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
	"github.com/tiandi/fpdb"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // parse arguments and put into r.Form
	fmt.Println("form: ", r.Form)
	fmt.Println("path: ", r.URL.Path)
	fmt.Println("scheme: ", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello footprint!")
}

func main() {
	dbConfig := mysql.Config{"fp:123@tcp(127.0.0.1:3306)/", 3, 3}
	db := mysql.NewMysql(&dbConfig)
	mysql.CreateNewDatabase(db, "testDB")
	mysql.CreateNewTable(db)
	http.HandleFunc("/", sayhelloName)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndserve: ", err)
	}
}
