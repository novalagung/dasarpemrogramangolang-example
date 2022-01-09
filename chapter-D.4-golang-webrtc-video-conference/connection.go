package main

import (
	"sync"

	"github.com/gorilla/websocket"
)

type WebSocketConnection struct {
	*websocket.Conn
	RoomID string
	UserID string
}

type WebSocketRoom struct {
	RoomID      string
	Connections map[string]*WebSocketConnection
}

type WebSocketManager struct {
	sync.Mutex
	Rooms map[string]*WebSocketRoom
}

var manager = WebSocketManager{
	Rooms: make(map[string]*WebSocketRoom),
}

func (m *WebSocketManager) AppendConnection(conn *WebSocketConnection) {
	m.Lock()
	defer m.Unlock()

	if _, isRoomExists := m.Rooms[conn.RoomID]; !isRoomExists {
		m.Rooms[conn.RoomID] = &WebSocketRoom{
			RoomID:      conn.RoomID,
			Connections: make(map[string]*WebSocketConnection),
		}
	}

	room := m.Rooms[conn.RoomID]
	if _, isConnectionExists := room.Connections[conn.UserID]; !isConnectionExists {
		room.Connections[conn.UserID] = conn
	}

}

func (m *WebSocketManager) EjectConnection(conn *WebSocketConnection) {
	m.Lock()
	defer m.Unlock()

	if _, isRoomExists := m.Rooms[conn.RoomID]; !isRoomExists {
		return
	}

	room := m.Rooms[conn.RoomID]
	delete(room.Connections, conn.UserID)
}

func (m *WebSocketManager) GetSiblingConnections(conn *WebSocketConnection) []*WebSocketConnection {
	m.Lock()
	defer m.Unlock()

	connections := make([]*WebSocketConnection, 0)

	if _, isRoomExists := m.Rooms[conn.RoomID]; !isRoomExists {
		return connections
	}

	for _, each := range m.Rooms[conn.RoomID].Connections {
		// if each.UserID == conn.UserID {
		// 	continue
		// }
		connections = append(connections, each)
	}

	return connections
}
