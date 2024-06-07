package main

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"log"
)

func SetupSocketIo() *socketio.Server {
	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()

	defer func(server *socketio.Server) {
		err := server.Close()
		if err != nil {
			log.Println(err)
		}
	}(server)

	return server
}
