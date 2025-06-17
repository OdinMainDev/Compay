package eswebsocket

import "sync"

type roomService struct {
	roomMap map[string]*room
	mutex   sync.Mutex
}

func (rs *roomService) initRoomService() {
	rs.roomMap = make(map[string]*room)
}

func (rs *roomService) createRoom(roomName string) *room {
	if room, ok := rs.getRoomByName(roomName); ok {
		return room
	}

	rs.mutex.Lock()
	defer rs.mutex.Unlock()

	room := &room{name: roomName}
	room.initRoomData()
	rs.roomMap[room.getRoomName()] = room
	return room
}

func (rs *roomService) getRoomByName(roomName string) (*room, bool) {
	rs.mutex.Lock()
	defer rs.mutex.Unlock()
	room, ok := rs.roomMap[roomName]
	return room, ok
}
