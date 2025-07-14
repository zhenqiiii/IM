package wbskt

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

// http升级WebSocket所用结构体
var upgrader = websocket.Upgrader{}

// 存放WebSocket连接
var ws = make(map[*websocket.Conn]struct{})

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	// 每次接收到请求，都将升级后的ws连接加入集合
	ws[c] = struct{}{}
	// 最外层的循环保持连接：持续读取连接中客户端发送的信息，并持续返回信息
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		// 循环发送给每个连接：实现一个人发送的消息所有在线的人都能接收到
		for conn := range ws {
			err = conn.WriteMessage(mt, message)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}

	}
}

// 检查端口是否被占用
func isPortAvailable(addr string) bool {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return false
	}
	_ = ln.Close()
	return true
}

// http库的WebSocket服务
func WebSocketServer() {
	if !isPortAvailable("localhost:8080") {
		log.Fatal("port 8080 is already in use")
	}
	fmt.Print("run websocket")
	http.HandleFunc("/echo", echo)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

// Gin框架的WebSocket服务
func GinWebSocketServer() {
	r := gin.Default()

	// router
	r.GET("/echo", func(ctx *gin.Context) {
		echo(ctx.Writer, ctx.Request)
	})

	r.Run(":8080")
}
