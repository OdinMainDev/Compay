package eswebsocket

type eventService struct {
	mainService *ESWebsocket
}

// func (es *eventService) getEvent(code int) string {
// 	return EventText(code)
// }

func (es *eventService) handleEvent(eventCode int, conn *Connection) {
	switch eventCode {
	case CONNECTION:
		es.handleConnection(conn)
		return
	case DISCONNECTION:
		es.handleDisconnection(conn)
		return
	}
}

func (es *eventService) handleConnection(conn *Connection) {
	es.mainService.connectionService.addConnection(conn)
}

func (es *eventService) handleDisconnection(conn *Connection) {
	conn.room.removeConnection(conn)
	es.mainService.connectionService.removeConnection(conn)
	conn.socketConnection.Close()
}
