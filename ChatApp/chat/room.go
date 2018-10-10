package main

import (
	"github.com/gorilla/websocket"
	"net/http"
	"log"
	"GoPrecices/ChatApp/tracer"
	"github.com/stretchr/objx"
)

const (
	socketBufferSize = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: messageBufferSize}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request){
	socket, err := upgrader.Upgrade(w, req,nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	authCookie, err := req.Cookie("auth")
	if err != nil{
		log.Fatal("Failed to get auth cookie: ",err)
		return
	}


	client := &client{
		socket: socket,
		send: make(chan *message, messageBufferSize),
		room: r,
		userData: objx.MustFromBase64(authCookie.Value),
	}

	r.join <- client
	defer func() {r.leave <- client}()
	go client.write()
	client.read()
}

type room struct {

	forward chan *message

	join chan *client

	leave chan *client

	clients map[*client]bool

	tracer tracer.Tracer

	avatar Avatar

}

func newRoom(avatar Avatar) *room {
	return &room{
		forward: make(chan *message),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map [*client]bool),
		tracer:  tracer.Off(),
}
}

func (r *room) run(){
	for {
		select {

		case client := <- r.join:
			r.clients[client] = true
			r.tracer.Trace("Client joined chat\n")

		case client := <- r.leave:
			delete(r.clients,client)
			close(client.send)
			r.tracer.Trace("Client left chat\n")

		case msg := <- r.forward:
			for client := range r.clients{
				client.send <- msg
				r.tracer.Trace("New message:[%s], sent to client\n	",msg.Message)
			}
		}


	}
}


