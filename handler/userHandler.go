package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"newim/entity"
	"newim/service/impl"
)

func UserRegisterHandler(c *gin.Context) {
	userId, pwd := c.PostForm("userId"), c.PostForm("pwd")
	userService := &impl.UserServiceImpl{}
	err := userService.UserRegister(userId, pwd)
	if err != nil {
		c.JSON(200, gin.H{
			"code": "1",
			"msg":  err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"code": "0",
			"msg":  "success",
		})
	}
}

func UserWebSocketHandler(c *gin.Context) {
	fromId, toId := c.Query("fromId"), c.Query("toId")

	conn, err := (&websocket.Upgrader{ // 建立websocket连接
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Fatal(err)
	}

	client := &entity.Client{
		FromID:     fromId,
		ToID:       toId,
		Conn:       conn,
		MsgChannel: make(chan []byte),
	}

	entity.Svr.Register(client)

	// 读
	go func(c *entity.Client) {
		entity.Svr.SendMsgToClient(c)
	}(client)

	// 写
	go func(c *entity.Client) {
		entity.Svr.GetMsgFromClient(c)
	}(client)

}
