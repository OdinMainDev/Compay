package eswebsocket

import (
	"compay/core/eslog"
	"sync"

	"github.com/gorilla/websocket"
)

type Connection struct {
	sessionId           string
	remoteAddress       string
	userAgent           string
	isClientInitialized bool
	customData          map[string]interface{}
	socketConnection    *websocket.Conn
	room                *room
	mutex               sync.Mutex
	wss                 *WebSocketService
}

func (conn *Connection) GetRemoteAddress() string {
	conn.mutex.Lock()
	defer conn.mutex.Unlock()
	return conn.remoteAddress
}

func (conn *Connection) GetUserAgent() string {
	conn.mutex.Lock()
	defer conn.mutex.Unlock()
	return conn.userAgent
}

func (conn *Connection) GetSessionId() string {
	conn.mutex.Lock()
	defer conn.mutex.Unlock()
	return conn.sessionId
}

func (conn *Connection) GetIsClientInitialized() bool {
	conn.mutex.Lock()
	defer conn.mutex.Unlock()
	return conn.isClientInitialized
}

func (conn *Connection) GetData(key string) (interface{}, bool) {
	conn.mutex.Lock()
	defer conn.mutex.Unlock()
	val, ok := conn.customData[key]
	return val, ok
}

func (conn *Connection) SetData(key string, val interface{}) {
	conn.mutex.Lock()
	defer conn.mutex.Unlock()
	conn.customData[key] = val
}

func (conn *Connection) SetInitialized(roomName string, sessionId string) {
	room, ok := GetInstance().getRoomService().getRoomByName(roomName)
	if !ok {
		room = GetInstance().getRoomService().createRoom(roomName)
	}

	conn.sessionId = sessionId
	conn.isClientInitialized = true

	room.addConnection(conn)
	GetInstance().WebSocketService.handleEvent(CONNECTION, conn)
}

func (conn *Connection) SendTextMessage(message []byte) {
	conn.mutex.Lock()
	defer conn.mutex.Unlock()
	conn.sendMessage(message, websocket.TextMessage)
}

func (conn *Connection) SendBinaryMessage(message []byte) {
	conn.mutex.Lock()
	defer conn.mutex.Unlock()
	conn.sendMessage(message, websocket.BinaryMessage)
}

func (conn *Connection) sendMessage(message []byte, protocolType int) {
	err := conn.socketConnection.WriteMessage(protocolType, message)
	if err != nil {
		eslog.LogError("ESWS SendMessageError: ", err.Error(), protocolType)
	}
}

func (conn *Connection) setRoom(r *room) {
	conn.mutex.Lock()
	defer conn.mutex.Unlock()
	conn.room = r
}

func (conn *Connection) getRoom() *room {
	conn.mutex.Lock()
	defer conn.mutex.Unlock()
	return conn.room
}

func (conn *Connection) KillConnection() {
	conn.wss.handleEvent(DISCONNECTION, conn)
}
