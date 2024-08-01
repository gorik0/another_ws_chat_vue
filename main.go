package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func main() {

	gi := gin.Default()
	gi.NoRoute(handler)
	err := http.ListenAndServe(":9000", gi)
	if err != nil {
		panic(err)
	}
}

func handler(context *gin.Context) {

	log.Println("upgrade connection")

	writer := context.Writer
	up := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	conn, err := up.Upgrade(writer, context.Request, nil)
	if err != nil {

		log.Println(err)
	}

	go handleConn(conn)

}

func handleConn(c *websocket.Conn) {
	go startPingPongSession(c)
	for {
		messageType, p, err := c.ReadMessage()
		if err != nil {
			log.Printf("error reading from websocket: %v", err)
			return
		}
		log.Printf("Got message type: %v, %v\n", messageType, p)
		sendACK(c)
	}
}

func startPingPongSession(c *websocket.Conn) {
	c.SetPongHandler(func(pong string) error {
		println("pong received with ", pong)
		return nil
	})
	c.SetPingHandler(func(ping string) error {
		println("png received with ", ping)
		return nil
	})
	ticker := time.NewTicker(3 * time.Second)
	for {
		select {
		case <-ticker.C:
			c.SetWriteDeadline(time.Now().Add(3 * time.Second))
			println("ping sending...")
			err := c.WriteMessage(websocket.PingMessage, []byte("ping"))
			if err != nil {
				fmt.Println("Error sending ping: ", err)
				return
			}

		}
	}
}

func sendACK(c *websocket.Conn) {
	ack := struct {
		Typ    string `json:"type"`
		Msg    string `json:"msg"`
		Status string `json:"status"`
	}{
		Typ:    "ACK",
		Status: "OK",
	}
	ackBytes, _ := json.Marshal(ack)
	fmt.Printf("ackBytes: %v\n", ackBytes)
	err := c.WriteJSON(ack)
	if err != nil {
		log.Println("Error sending ACK: ", err)
		return
	}
}
