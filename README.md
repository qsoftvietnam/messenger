# Go - Web Messenger

This is a simple chat web app written in Go

INTRODUCE
--------------------------------------------------------------------------------
Web Messenger
Develop a web application and backend for chatting.
```
1. User asked for a display name before entering the chatroom.
2. Single chat room for all users.
3. User able to send text into the chatroom.
4. Chat history should be saved for the last 100 messages. Every user should see the chat
history once they are in the chatroom.
```

INSTALL
--------------------------------------------------------------------------------
Just run the following

```
1. download & install go: https://golang.org/dl/
2. create project structure
    - project (Your project)
        - bin
        - pkg        
        - src
        - public  

3. config GOPATH in 'Environment Variables' = Directory (Your project)
4. cmd: go get github.com/qsoftvietnam/messenger
5. install packages
    - go get github.com/go-sql-driver/mysql
    - go get github.com/gorilla/websocket
    - go get github.com/icza/session

6. sql: excecute file chat.sql in $GOPATH/src/github.com/qsoftvietnam/messenger/database
7. cmd: cd  $GOPATH/src/github.com/qsoftvietnam/messenger/src
8. config the environment variables in main.go
    - os.Setenv("port", ":8080")
    - os.Setenv("dbPort", "3306")
    - os.Setenv("dbHost", "127.0.0.1")
    - os.Setenv("dbUsername", "root")
    - os.Setenv("dbPassword", "")
    - os.Setenv("dbName", "chat")
9. go build main.go handles.go models.go runtimes.go structs.go OR go build *.go
10. cmd: main.exe OR "filename".exe
```

Then point your browser to http://localhost:8080

TECHNOLOGY
--------------------------------------------------------------------------------

```
- Chat window:​ Web technologies (HTML, javascript, CSS, VueJs, Material, etc.)
- Backend: GoLang
- Database: Mysql
```
