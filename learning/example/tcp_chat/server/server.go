package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

type Server struct {
	Protocol string
	Addr     string

	Listener net.Listener

	joinRoomCh  chan string
	leaveRoomCh chan string

	// room_name: room
	Room map[string]*Room
	// client_id: client
	Clients map[string]*Client

	m sync.Mutex
}

func NewServer(proto, addr string) *Server {
	listener, err := net.Listen(proto, addr)
	if err != nil {
		log.Fatal(err)
	}

	return &Server{
		Protocol: proto,
		Addr:     addr,
		Listener: listener,

		joinRoomCh: make(chan string),

		Room:    make(map[string]*Room),
		Clients: make(map[string]*Client),
	}
}

func (s *Server) Listen() {
	log.Println("[+] Listening...")
	for {
		client, err := s.Listener.Accept()
		if err != nil {
			log.Println("Error when accept new client", err)
			continue
		}

		go s.HandleNewClient(&client)
	}
}

func (s *Server) HandleNewClient(clientConn *net.Conn) {
	// read client info for the first time
	reader := bufio.NewReader(*clientConn)
	info, err := reader.ReadString('\n')
	if err != nil {
		log.Println("client connect err: ", err)
		return
	}

	log.Println("[+] Receive new client connect:", info)

	clientId, roomName := extractClientInfo(info)
	s.m.Lock()
	// disconnect if duplicate client_id
	if _, exists := s.Clients[clientId]; exists {
		m := fmt.Sprintf("Client %v already logged in. Please try a new name !!!", clientId)
		_, err := (*clientConn).Write([]byte(m))
		if err != nil {
			log.Println(err)
		}
		log.Printf("Client %v already logged in! Disconect...", clientId)
		time.Sleep(time.Second * 1)
		(*clientConn).Close()
		return
	}

	// save client
	client := NewClient(clientId, roomName, clientConn)
	s.Clients[clientId] = client
	s.m.Unlock()

	// create new room if not exists
	s.m.Lock()
	if _, exists := s.Room[roomName]; !exists {
		room := NewRoom(roomName)
		s.Room[roomName] = room
	}
	s.m.Unlock()

	s.Room[roomName].addClient(clientId, client)

	// send message to this room
	joinMsg := fmt.Sprintf("[+] New client: %v\n", clientId)
	s.Room[roomName].Publish(clientId, joinMsg)

	// start wait new message from client
	for {
		reader = bufio.NewReader(*clientConn)
		newMsg, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("[+] Client %v disconnected !!! %v\n", clientId, err)
			break
		}

		newMsg = fmt.Sprintf("%v: %v", clientId, newMsg)
		s.Room[roomName].Publish(clientId, newMsg)
	}

	// clear client if disconnected
	leaveMsg := fmt.Sprintf("[+] Client has leave this room: %v\n", clientId)
	s.Room[roomName].Publish(clientId, leaveMsg)

	s.Room[roomName].rmClient(clientId)
	s.rmClient(clientId)

	log.Printf("[+] Client %v has leave room %v!!!\n", clientId, roomName)
}

func (s *Server) rmClient(id string) {
	if _, exists := s.Clients[id]; !exists {
		return
	}

	s.m.Lock()
	delete(s.Clients, id)
	s.m.Unlock()
}

// client_id:Room
func extractClientInfo(s string) (clientId, room string) {
	lst := strings.SplitN(s, ":", 2)
	clientId, room = lst[0], lst[1]
	room = room[:len(room)-1]
	log.Println("[debug] client: ", clientId)
	log.Println("[debug] room: ", room)
	return
}
