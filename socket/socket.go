package socket

import (
	"contact/model"
	"fmt"
	"log"

	"github.com/17twenty/shortuuid"
	guuid "github.com/google/uuid"
	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)

var Server *gosocketio.Server

func init() {
	Server = gosocketio.NewServer(transport.GetDefaultWebsocketTransport())
	LoadSocket()
	fmt.Println("Socket Inititalize...")
}

func LoadSocket() {
	Server.On("/create", func(c *gosocketio.Channel) string {
		fmt.Println("create called")
		// c.Emit("/message", Message{"using emit"})
		_ = guuid.New().String()
		id := shortuuid.New()
		fmt.Println(c.List(id))
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
	Server.On("/join", func(c *gosocketio.Channel, channel model.Channel) string {
		fmt.Println("receiving join request")
		if len(c.List(channel.Channel)) == 0 {
			return "Need to create a channel " + channel.Channel
		}
		log.Println("Client joined to ", channel.Channel)
		c.Join(channel.Channel)
		// c.BroadcastTo(channel.Channel, "/message", Message{"using broadcast"})
		return "joined to " + channel.Channel
	})

	Server.On("/message", func(c *gosocketio.Channel, args model.Message) {
		fmt.Println("receiving message from server", args)
		// chanee, _ := server.GetChannel(c.Id())
		// fmt.Println(chanee, c.Id(), c)
		c.BroadcastTo(args.Channel, "/message", args)
	})

	Server.On(gosocketio.OnDisconnection, func(c *gosocketio.Channel) {
		log.Println("Disconnected")
	})
}
