package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/17twenty/shortuuid"
	"github.com/gorilla/mux"

	guuid "github.com/google/uuid"
	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)

type Channel struct {
	Channel string `json:"channel"`
}

type Message struct {
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

var router *mux.Router

func main() {
	router = mux.NewRouter()
	router.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./asset"))))
	server := gosocketio.NewServer(transport.GetDefaultWebsocketTransport())
	// server.On(gosocketio.OnConnection, func(c *gosocketio.Channel) {
	// log.Println("Created lobby", counter)

	// c.Emit("/message", Message{10, "main", "using emit"})
	// var id string
	// id = str[counter]
	// c.Join(id)
	// c.BroadcastTo(id, "/message", Message{10, "main", "using broadcast"})
	// counter++
	// })

	server.On("/create", func(c *gosocketio.Channel) string {
		fmt.Println("create called")
		// c.Emit("/message", Message{"using emit"})
		_ = guuid.New().String()
		id := shortuuid.New()
		if len(c.List(id)) > 0 {
			fmt.Println("Error creating code")
			return "Error Creating Code"
		}
		fmt.Println(id)
		c.Join(id)
		log.Println("Created lobby", id)
		fmt.Println(c.List(id))
		// c.BroadcastTo(id, "/message", Message{"using broadcast"})
		return id
	})
	server.On("/join", func(c *gosocketio.Channel, channel Channel) string {
		fmt.Println("receiving join request")
		if len(c.List(channel.Channel)) == 0 {
			return "Need to create a channel " + channel.Channel
		}
		log.Println("Client joined to ", channel.Channel)
		c.Join(channel.Channel)
		// c.BroadcastTo(channel.Channel, "/message", Message{"using broadcast"})
		return "joined to " + channel.Channel
	})

	server.On("/message", func(c *gosocketio.Channel, args Message) {
		fmt.Println("receiving message from server", args)
		// chanee, _ := server.GetChannel(c.Id())
		// fmt.Println(chanee, c.Id(), c)
		c.BroadcastTo(args.Channel, "/message", args)
	})

	server.On(gosocketio.OnDisconnection, func(c *gosocketio.Channel) {
		log.Println("Disconnected")
	})

	router.Handle("/socket.io/", server)

	log.Fatal(http.ListenAndServe(":8080", router))

	fmt.Println("")
}

// "/message/id"
