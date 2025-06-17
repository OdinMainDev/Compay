package eswebsocket

import (
	"compay/core/eslog"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type WebSocketService struct {
	eventService *eventService
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var (
	pongWait = 20 * time.Second
)

type ListenHelper struct {
	Writer           http.ResponseWriter
	Request          *http.Request
	UserIp           string
	UserAgent        string
	HandlerFunc      func([]byte, *Connection) error
	AutoInitFunc     func(*Connection, ...interface{}) error
	AutoInitFuncArgs interface{}
}

func (wss *WebSocketService) Listen(data ListenHelper) {
	c, err := upgrader.Upgrade(data.Writer, data.Request, nil)
	if err != nil {
		eslog.LogError(err)
		return
	}

	conn := &Connection{
		remoteAddress:    data.UserIp,
		userAgent:        data.UserAgent,
		socketConnection: c,
		customData:       make(map[string]interface{}),
		wss:              wss,
	}

	if data.AutoInitFunc != nil {
		err := data.AutoInitFunc(conn, data.AutoInitFuncArgs)
		if err != nil {
			eslog.LogError("ESWS AutoInitError: ", err.Error())
			wss.handleEvent(DISCONNECTION, conn)
		}
	}
	defer c.Close()
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			wss.handleEvent(DISCONNECTION, conn)
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				eslog.LogError("ESWS UnexpectedCloseError: ", err.Error())
			}
			break
		}

		if err := data.HandlerFunc(message, conn); err != nil {
			eslog.LogError("ESWS HandlerError: ", err.Error())
			wss.handleEvent(DISCONNECTION, conn)
			break
		}

		conn.socketConnection.SetReadDeadline(time.Now().Add(pongWait))
	}
}

func (wss *WebSocketService) handleEvent(eventCode int, conn *Connection) {
	wss.eventService.handleEvent(eventCode, conn)
}
