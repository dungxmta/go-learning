package main

import (
	"log"
	"sync"
)

type Room struct {
	Name    string
	Clients map[string]*Client
	// Clients map[*net.Conn]bool
	m sync.Mutex
}

func NewRoom(name string) *Room {
	return &Room{
		Name:    name,
		Clients: make(map[string]*Client),
		// Clients: make(map[*net.Conn]bool),
	}
}

func (r *Room) Publish(sender string, msg string) {
	for clientId, client := range r.Clients {
		if client == nil {
			delete(r.Clients, clientId)
			continue
		}

		// skip sender
		if sender == clientId {
			continue
		}

		err := client.Publish(msg)
		if err != nil {
			log.Printf("[Error] Can't send msg to client %v: %v\n", clientId, err)
		}
	}
}

func (r *Room) addClient(id string, client *Client) bool {
	if _, exists := r.Clients[id]; exists {
		return false
	}

	r.m.Lock()
	r.Clients[id] = client
	r.m.Unlock()

	return true
}

func (r *Room) rmClient(id string) {
	if _, exists := r.Clients[id]; !exists {
		return
	}

	r.m.Lock()
	delete(r.Clients, id)
	r.m.Unlock()
}
