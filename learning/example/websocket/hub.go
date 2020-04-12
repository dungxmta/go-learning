package main

type Message struct {
	data []byte
	room string
}

type Subscription struct {
	conn *Connection
	room string
}

// Hub maintains the set of active connections and broadcasts messages to the
// connections.
type Hub struct {
	// Registered connections.
	rooms map[string]map[*Connection]bool

	// Inbound messages from the connections.
	broadcast chan Message

	// Register requests from the connections.
	register chan Subscription

	// Unregister requests from connections.
	unregister chan Subscription
}

var h = Hub{
	broadcast:  make(chan Message),
	register:   make(chan Subscription),
	unregister: make(chan Subscription),
	rooms:      make(map[string]map[*Connection]bool),
}

func (h *Hub) Run() {
	for {
		select {
		case s := <-h.register:
			connections := h.rooms[s.room]
			if connections == nil {
				connections = make(map[*Connection]bool)
				h.rooms[s.room] = connections
			}
			h.rooms[s.room][s.conn] = true
		case s := <-h.unregister:
			connections := h.rooms[s.room]
			if connections != nil {
				if _, ok := connections[s.conn]; ok {
					delete(connections, s.conn)
					close(s.conn.send)
					if len(connections) == 0 {
						delete(h.rooms, s.room)
					}
				}
			}
		case m := <-h.broadcast:
			connections := h.rooms[m.room]
			for c := range connections {
				select {
				case c.send <- m.data:
				default:
					close(c.send)
					delete(connections, c)
					if len(connections) == 0 {
						delete(h.rooms, m.room)
					}
				}
			}
		}
	}
}
