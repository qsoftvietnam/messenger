package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/websocket"
	"github.com/icza/session"
)

const timeFormat = "20060102150405"

var db *sql.DB
var err error
var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Info)              // broadcast channel
var upgrader = websocket.Upgrader{           // Configure the upgrader
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	// Config environment
	os.Setenv("port", ":8080")
	os.Setenv("dbPort", "3306")
	os.Setenv("dbHost", "127.0.0.1")
	os.Setenv("dbUsername", "root")
	os.Setenv("dbPassword", "")
	os.Setenv("dbName", "chat")

	// For testing purposes, we want cookies to be sent over HTTP too (not just HTTPS):
	session.Global.Close()
	session.Global = session.NewCookieManagerOptions(session.NewInMemStore(), &session.CookieMngrOptions{AllowHTTP: true})

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err = sql.Open("mysql", os.Getenv("dbUsername")+":@tcp("+os.Getenv("dbHost")+":"+os.Getenv("dbPort")+")/"+os.Getenv("dbName"))

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// Create a simple file server
	fs := http.FileServer(http.Dir("../public"))
	http.Handle("/", fs)

	// Configure websocket route
	http.HandleFunc("/ws", handleConnections)

	// Configure join handle
	http.HandleFunc("/join", handleJoin)

	// Start listening for incoming chat messages
	go handleMessages()

	// Start the server on localhost port 8080 and log any errors
	log.Println("http server started on ", os.Getenv("port"))
	err := http.ListenAndServe(os.Getenv("port"), nil)
	// err := http.ListenAndServe(":8080", context.ClearHandler(http.DefaultServeMux))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
