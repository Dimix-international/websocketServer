package websocket

import "fmt"

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (p *Pool) Start() {
	for {
		select {
		case client := <-p.Register:
			p.Clients[client] = true
			fmt.Println("size of connection pull: ", len(p.Clients))
			for client := range p.Clients {
				fmt.Println(client)
				client.Conn.WriteJSON(Message{Type: 1, Body: "New user joined..."})
			}
		case client := <- p.Unregister:
			delete(p.Clients, client)
			fmt.Println("size of connection pull after disconnect: ", len(p.Clients))
			for client := range p.Clients {
				fmt.Println(client)
				client.Conn.WriteJSON(Message{Type: 1, Body: "User disconnected..."})
			}
		case message := <- p.Broadcast:
			fmt.Println("Sending message to all clients in the pool")
			for client := range p.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}