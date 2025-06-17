package eswebsocket

import (
	"compay/core/eslog"
	"fmt"
	"sync"
	"time"
)

type ESWebsocket struct {
	connectionService *connectionService
	eventService      *eventService
	roomService       *roomService
	WebSocketService  *WebSocketService
	mutex             sync.Mutex
}

var eswsInstance *ESWebsocket
var lock = &sync.Mutex{}

func GetInstance() *ESWebsocket {
	if eswsInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if eswsInstance == nil {
			eswsInstance = &ESWebsocket{}
			eswsInstance.Init()
		}
	}
	return eswsInstance
}

func (esws *ESWebsocket) Init() {
	esws.eventService = &eventService{mainService: esws}
	esws.connectionService = &connectionService{}
	esws.connectionService.Init()

	esws.WebSocketService = &WebSocketService{
		eventService: esws.eventService,
	}

	esws.roomService = &roomService{}
	esws.roomService.initRoomService()
	if 1 == 0 {
		go esws.runTest()
	}
}

func (esws *ESWebsocket) GetConnectionBySessionId(sessionId string) (*Connection, bool) {
	conn, ok := esws.getConnectionService().getConnectionsMap()[sessionId]
	return conn, ok
}

func (esws *ESWebsocket) CreateRoom(roomName string) {
	esws.mutex.Lock()
	defer esws.mutex.Unlock()
	esws.getRoomService().createRoom(roomName)
}

func (esws *ESWebsocket) getConnectionService() *connectionService {
	return esws.connectionService
}

func (esws *ESWebsocket) getRoomService() *roomService {
	return esws.roomService
}

func (esws *ESWebsocket) SendBinaryMessageByRoomName(roomName string, message []byte) {
	room, ok := esws.getRoomService().getRoomByName(roomName)
	if !ok {
		eslog.LogError("ESWS BadRoom: ", roomName)
		return
	}
	room.sendBinaryMessageToUsers(message)
}

func (esws *ESWebsocket) SendTextMessageByRoomName(roomName string, message []byte) {
	room, ok := esws.getRoomService().getRoomByName(roomName)
	if !ok {
		eslog.LogError("ESWS BadRoom: ", roomName)
		return
	}
	room.sendTextMessageToUsers(message)
}

func (esws *ESWebsocket) GetWebsoketService() *WebSocketService {
	return esws.WebSocketService
}

func (esws *ESWebsocket) GetRoomByName(roomName string) (*room, bool) {
	return esws.getRoomService().getRoomByName(roomName)
}

func (esws *ESWebsocket) SendBinaryMessageBySessionId(sessionName string, message []byte) {
	conn, ok := esws.GetConnectionBySessionId(sessionName)
	if !ok {
		return
	}
	conn.SendBinaryMessage(message)
}

func (esws *ESWebsocket) runTest() {
	for {
		fmt.Println("CONNECTIONS: ", esws.connectionService.getConnectionsCount())
		fmt.Println("ROOMS: ", len(esws.roomService.roomMap))
		time.Sleep(time.Millisecond * 500)
	}
}
