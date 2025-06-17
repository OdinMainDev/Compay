package eswebsocket

const (
	CONNECTION    int = 1
	DISCONNECTION int = 2
)

var EventsText = map[int]string{
	CONNECTION:    "Connection",
	DISCONNECTION: "Disconnection",
}

func EventText(code int) string {
	return EventsText[code]
}
