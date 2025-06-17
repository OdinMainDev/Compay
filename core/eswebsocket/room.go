package eswebsocket

import "sync"

type room struct {
	name           string
	connectionsMap map[string]*Connection
	mutex          sync.Mutex
}

func (r *room) initRoomData() {
	r.connectionsMap = make(map[string]*Connection)
}

func (r *room) getRoomName() string {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.name
}

func (r *room) addConnection(conn *Connection) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	sessionId := conn.GetSessionId()
	r.connectionsMap[sessionId] = conn
	conn.setRoom(r)
}

func (r *room) removeConnection(conn *Connection) {
	if conn == nil || r == nil {
		return
	}
	r.mutex.Lock()
	defer r.mutex.Unlock()
	delete(r.connectionsMap, conn.GetSessionId())
}

func (r *room) sendBinaryMessageToUsers(message []byte) {
	connections := r.getConnectionsMap()
	r.mutex.Lock()
	defer r.mutex.Unlock()
	for _, conn := range connections {
		conn.SendBinaryMessage(message)
	}
}

func (r *room) sendTextMessageToUsers(message []byte) {
	connections := r.getConnectionsMap()
	r.mutex.Lock()
	defer r.mutex.Unlock()
	for _, conn := range connections {
		conn.SendTextMessage(message)
	}
}

func (r *room) getConnectionsMap() map[string]*Connection {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.connectionsMap
}
