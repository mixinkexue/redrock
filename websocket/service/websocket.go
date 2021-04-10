package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)
type Client struct {
	User    string
	Conn    *websocket.Conn
	MsgChan chan Message
}
type Message struct {
	User    string `json:"user"`
	Content string `json:"content"`
}
type Manager struct {
	Chans map[string]chan Message
}
func (c *Client) Recv() {
	for {
		_, bytes, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		content := string(bytes)
		msg := Message{
			User:    c.User,
			Content: content,
		}
		manager.Broad(msg)
	}
}
func (m *Manager) Broad(msg Message) {
	for _, ch := range m.Chans {
		ch <- msg
	}
}
func (c *Client) Write()  {
	for {
		select {
		case m := <-c.MsgChan:
			err := c.Conn.WriteJSON(m)
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}
var manager Manager
var managers map[string]Manager
func Ws(ctx *gin.Context){
	room:=ctx.Param("room")
	managers=make(map[string]Manager)
	var ok bool
	manager,ok=managers[room]
	if !ok{
			manager := Manager{
				Chans: make(map[string]chan Message),
			}
			managers[room]=manager
	}
	upGrader := websocket.Upgrader{
		Subprotocols: []string{ctx.GetHeader("Sec-WebSocket-Protocol")},
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println(err)
		ctx.JSON(200, gin.H{
			"status": 10022,
			"info": "failed",
		})
		return
	}
	username, _ :=ctx.Cookie("username")
	client:=Client{
		User:    username,
		Conn:    conn,
		MsgChan: make(chan Message),
	}
	manager.Chans[client.User] = client.MsgChan
	go client.Recv()
	go client.Write()
}