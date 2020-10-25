# TL;DR

- Simple tcp app chat with multiple room
- Server -/> Multi room -/> User join room by name

## Logic

- When one client try to connect to server, 
 it sent message with format: client_id:room_name
- Server catch connect and split "client_id" & "room_name"
- Join client to room "room_name", so everytime client_id type a new message,
 it sent to server and publish to everyone in same room

## TODO

- Handle disconnect from client/server side
- Support command: quit, join_room, ...
- ...

## Demo

- server
```
cd ./server
go run .
```

- client 1
```
cd ./client
go run .
---
Enter NAME: A
Enter ROOM: Demo
```

- client 2
```
cd ./client
go run .
---
Enter NAME: B
Enter ROOM: Demo
```
