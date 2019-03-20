package main

import (
	"log"
	"net/http"
	"os"

	"github.com/alexbotello/simpleServer/trace"
	"github.com/gorilla/websocket"
)

type room struct {
	// forward is a channel that holds incoming messages that should be forwarded to the other clients.
	forward chan []byte
	join    chan *client
	leave   chan *client
	clients map[*client]bool
	tracer  trace.Off()
}

func newRoom() *room {
	// return a room pointer
	return &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
		tracer:  trace.New(os.Stdout),
	}
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
			r.tracer.Trace("New client joined")
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
			r.tracer.Trace("Client left")
		case msg := <-r.forward:
			r.tracer.Trace("Message received: ", string(msg))
			for client := range r.clients {
				client.send <- msg
				r.tracer.Trace(" -- sent to client")
			}
		}
	}
}

const (
	socketBuffersize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBuffersize, WriteBufferSize: socketBuffersize}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}
	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   r,
	}
	r.join <- client
	defer func() { r.leave <- client }()
	go client.write()
	client.read()
}
