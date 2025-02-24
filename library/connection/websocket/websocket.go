package websocket

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/websocket"
)


func New(w http.ResponseWriter, request *http.Request, logger *slog.Logger) *websocket.Conn {

	upgrader := websocket.Upgrader{}

	connection, err := upgrader.Upgrade(w, request, nil)
	
	if err != nil {
		logger.Error(err.Error())
		return nil
	}
	
	return connection

}


func OnReceive(connection *websocket.Conn, callback func(message string) error, logger *slog.Logger) bool {

	is_success := true

	_, message, err := connection.ReadMessage()

	if err != nil {
		logger.Error(err.Error())
		is_success = false
		return is_success
	}

	callback_err := callback(string(message))

	if callback_err != nil {
		logger.Error(callback_err.Error())
		is_success = false
	}

	return is_success

}


func Send(connection *websocket.Conn, message string, logger *slog.Logger) bool {

	is_success := true

	err := connection.WriteMessage(websocket.TextMessage, []byte(message))

	if err != nil {
		logger.Error(err.Error())
		is_success = false
	}

	return is_success

}