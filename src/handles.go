package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/icza/session"
)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Get a session. We're ignoring the error resulted from decoding an
	// existing session: Get() always returns a session, even if empty.
	sess := session.Get(r)

	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Make sure we close the connection when the function returns
	defer ws.Close()
	// log.Println(clients)
	// Register our new client
	clients[ws] = true

	for {
		var msg Info
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}

		if msg.Message != "" {
			if sess != nil {
				message := Message{}
				message.UserID = sess.CAttr("userID").(int)
				message.Message = msg.Message
				message.CreatedAt = time.Now().Format(timeFormat)

				go insertMessage(message)

				broadcast <- msg
			} else {
				log.Println("Session is nil")
			}
		} else {
			var userID int
			checkUser, _ := getUser(msg.Username) // Check existing by username

			if checkUser == (User{}) {
				if msg.Username != "" && msg.Email != "" {
					user := User{}
					user.Username = msg.Username
					user.Email = msg.Email
					user.CreatedAt = time.Now().Format(timeFormat)

					lastID, _ := insertUser(user)
					userID = int(lastID)
				}
			} else {
				userID = int(checkUser.ID)
			}

			if sess == nil {
				// Successful login. New session with initial constant and variable attributes:
				if userID > 0 {
					sess = session.NewSessionOptions(&session.SessOptions{
						CAttrs: map[string]interface{}{"userID": userID},
					})

					session.Add(sess, w)
				}
			}
		}
	}
}

func handleJoin(w http.ResponseWriter, r *http.Request) {
	messages, errGetMessages := getMessages()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if errGetMessages != nil {
		log.Printf("error: %v", errGetMessages)

		json.NewEncoder(w).Encode(false)
	} else {
		json.NewEncoder(w).Encode(messages)
	}
}
