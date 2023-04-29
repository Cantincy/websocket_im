package entity

import (
	"bytes"
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

type Client struct {
	FromID     string
	ToID       string
	Conn       *websocket.Conn
	MsgChannel chan []byte
}

type Msg struct {
	FromID string
	ToID   string
	Msg    []byte
}

type Server struct {
	ClientMap map[string]*Client
	MsgQueue  chan *Msg
	RWMutex   sync.RWMutex
}

var Svr = &Server{
	ClientMap: make(map[string]*Client),
	MsgQueue:  make(chan *Msg),
}

func (s *Server) DispatchMsg() { // server持续向接收方转发消息
	for {
		msg := <-s.MsgQueue
		toClient := s.ClientMap[msg.ToID]
		if toClient != nil {
			toClient.MsgChannel <- msg.Msg
			log.Printf("[Server]向client%s传递消息...", msg.ToID)
		}
	}
}

func (s *Server) Register(c *Client) { // server注册客户端消息（用户上线）
	s.RWMutex.Lock()
	s.ClientMap[c.FromID] = c
	s.RWMutex.Unlock()
}

func (s *Server) UnRegister(c *Client) { // server删除客户端消息（用户下线）
	defer c.Conn.Close()
	s.RWMutex.Lock()
	delete(s.ClientMap, c.FromID)
	s.RWMutex.Unlock()
}

func (s *Server) SendMsgToClient(c *Client) { // server向客户端发送数据
	for {
		msg := <-c.MsgChannel
		log.Printf("[Client%s]接收消息...", c.FromID)
		c.Conn.WriteMessage(websocket.TextMessage, msg)
	}
}

func (s *Server) GetMsgFromClient(c *Client) { // server从客户端接收数据
	for {
		_, buf, err := c.Conn.ReadMessage()
		if err != nil {
			return
		}

		if bytes.Compare(buf, []byte("exit")) == 0 {
			Svr.UnRegister(c)
			log.Printf("[Server]:Client%s下线", c.FromID)
			return
		}

		if Svr.ClientMap[c.ToID] == nil {
			c.Conn.WriteMessage(websocket.TextMessage, []byte("接收方不在线..."))
		} else {
			msg := &Msg{FromID: c.FromID, ToID: c.ToID, Msg: buf}
			Svr.MsgQueue <- msg
		}
	}
}
