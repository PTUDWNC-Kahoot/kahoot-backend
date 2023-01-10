package v1

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/spf13/cast"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleWs(hub *Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			panic(err)
		}

		presentationID := cast.ToUint32(c.Param("id"))

		client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256), roomID: presentationID}
		client.hub.register <- client

		// Allow collection of memory referenced by the caller by doing all work in
		// new goroutines.
		go client.writePump()
		go client.readPump()

		// send initalization data of present

		data := hub.presentations[client.roomID]
		if data != nil {
			rawData, err := json.Marshal(Message{
				Action:  "initialize",
				Payload: data,
			})

			if err != nil {
				log.Println("failed to marshal data", err)
				return
			}

			client.send <- rawData
		}
	}
}