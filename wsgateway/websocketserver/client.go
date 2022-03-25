/*
Package websocketserver
@Time : 2022/3/25 11:06
@Author : fengxinlei
@File : client.go
@Software: GoLand
*/
package websocketserver

import (
	"fmt"
	"github.com/gorilla/websocket"
	"runtime/debug"
)

type Client struct {
	Addr          string
	Socket        *websocket.Conn
	Send          chan []byte
	AppId         uint32
	UserId        string
	FirstTime     uint64
	HeartbeatTime uint64
	LoginTime     uint64
}

func NewClient(addr string, socket *websocket.Conn, firstTime uint64) (client *Client) {
	return &Client{
		Addr:          addr,
		Socket:        socket,
		Send:          make(chan []byte, 100),
		FirstTime:     firstTime,
		HeartbeatTime: firstTime,
	}
}

func (c *Client) Read() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("write stop", string(debug.Stack()), r)
		}
	}()

	defer func() {
		fmt.Println("读取客户端数据 关闭send", c)
		close(c.Send)
	}()

	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			fmt.Println("读取客户端数据 错误", c.Addr, err)

			return
		}
		fmt.Println("读取客户端数据 处理:", string(message))
	}

}

func (c *Client) Write() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("write stop", string(debug.Stack()), r)
		}
	}()

	defer func() {
		c.Socket.Close()
		fmt.Println("Client发送数据 defer", c)
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				// 发送数据错误 关闭连接
				fmt.Println("Client发送数据 关闭连接", c.Addr, "ok", ok)

				return
			}
			c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}
