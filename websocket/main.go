
package main

import (
	"websocket/cmd"
	"websocket/dao"
)

func main()  {
	dao.MysqlInit()
	cmd.Entrance()
}
/*type Manager struct {
	Chans map[string]chan Message
}

func (m *Manager) Broad(msg Message) {
	for _, ch := range m.Chans {
		ch <- msg
	}
}

type Client struct {
	User    string
	Conn    *websocket.Conn
	MsgChan chan Message
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

type Message struct {
	User    string `json:"user"`
	Content string `json:"content"`
}

var manager Manager
var cnt = 0
*/
/*func main() {
	manager = Manager{
		Chans: make(map[string]chan Message),
	}
	r := gin.Default()

	r.GET("/ws", func(c *gin.Context) {
		upgrader := websocket.Upgrader{
			Subprotocols: []string{c.GetHeader("Sec-WebSocket-Protocol")},
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println(err)
			c.JSON(200, gin.H{
				"status": 10022,
				"info": "failed",
			})
			return
		}

		cnt++
		client := Client{
			Conn:    conn,
			MsgChan: make(chan Message),
			User:    strconv.Itoa(cnt),
		}
		manager.Chans[client.User] = client.MsgChan
		go client.Recv()
		go client.Write()
	})

	r.Run("127.0.0.1:8080")
}*/