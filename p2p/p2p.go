package p2p

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/hjkimGithub/nomadcoin/utils"
)

var upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(rw, r, nil)
	utils.HandleErr(err)

	for {
		_, p, err := conn.ReadMessage()
		// utils.HandleErr(err)
		if err != nil {
			// conn.Close()
			break
		}
		fmt.Printf("%s\n\n", p)
		message := fmt.Sprintf("New message: %s", p)
		utils.HandleErr(conn.WriteMessage(websocket.TextMessage, []byte(message)))
	}
}
