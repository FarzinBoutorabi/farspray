package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:82320591@/farspray")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	router := mux.NewRouter()
	router.Handle("/", http.FileServer(http.Dir("../Web/views/")))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../Web/static/"))))
	router.Handle("/signup", signUp).Methods("GET")
	//err = http.ListenAndServe(":80", handlers.LoggingHandler(os.Stdout, firewall(router)))
	err = http.ListenAndServeTLS(":443", "C:\\Users\\Hosein\\Documents\\Keys\\server.crt", "C:\\Users\\Hosein\\Documents\\Keys\\sample_key.priv", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
