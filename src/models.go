package main

import (
	"database/sql"
	"log"
)

// getUser - get a user by username
func getUser(username string) (User, error) {
	var user User
	// Execute the query
	err := db.QueryRow("SELECT * FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No user with that ID.")
	case err != nil:
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	return user, err
}

// insertUser - add new a user
func insertUser(data User) (int64, error) {
	// perform a db.Query insert
	insert, err := db.Exec("INSERT INTO users (username, email, created_at) VALUES (?, ?, ?)", data.Username, data.Email, data.CreatedAt)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	id, err := insert.LastInsertId()

	if err != nil {
		panic(err.Error())
	}

	return id, err
}

// getMessages - get 100 messages join users
func getMessages() (Infos, error) {
	// Execute the query
	rows, err := db.Query("SELECT username, email, message FROM messages INNER JOIN users ON users.id = messages.user_id ORDER BY messages.created_at asc LIMIT 100")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var messages Infos
	var message Info

	for rows.Next() {
		// for each row, scan the result into our tag composite object
		err = rows.Scan(&message.Username, &message.Email, &message.Message)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		messages = append(messages, message)
	}

	return messages, err
}

// insertMessage - add new a message
func insertMessage(data Message) {
	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO messages (user_id, message, created_at) VALUES (?, ?, ?)", data.UserID, data.Message, data.CreatedAt)
	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
}
