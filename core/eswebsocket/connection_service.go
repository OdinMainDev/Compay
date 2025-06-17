package eswebsocket

import "sync"

type connectionService struct {
	connectionMap map[string]*Connection
	mutex         sync.Mutex
}

func (cs *connectionService) Init() {
	cs.connectionMap = make(map[string]*Connection)
}

func (cs *connectionService) addConnection(conn *Connection) {
	cs.mutex.Lock()
	defer cs.mutex.Unlock()
	cs.connectionMap[conn.GetSessionId()] = conn
}

func (cs *connectionService) removeConnection(conn *Connection) {
	cs.mutex.Lock()
	defer cs.mutex.Unlock()
	delete(cs.connectionMap, conn.GetSessionId())
}

func (cs *connectionService) getConnectionsCount() int {
	cs.mutex.Lock()
	defer cs.mutex.Unlock()
	return len(cs.connectionMap)
}

func (cs *connectionService) getConnectionsMap() map[string]*Connection {
	cs.mutex.Lock()
	defer cs.mutex.Unlock()
	return cs.connectionMap
}
