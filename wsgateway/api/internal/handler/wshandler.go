package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"go-zero-ws/wsgateway/api/internal/svc"
	"go-zero-ws/wsgateway/websocketserver"
)

func wsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//l := logic.NewWsLogic(r.Context(), svcCtx)
		//err := l.Ws()
		//if err != nil {
		//	httpx.Error(w, err)
		//} else {
		//	httpx.Ok(w)
		//}
		ws := &websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				fmt.Println("升级协议", "ua:", r.Header["User-Agent"], "referer:", r.Header["Referer"])
				return true
			},
		}
		conn, err := ws.Upgrade(w, r, nil)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		fmt.Println("webSocket 建立连接:", conn.RemoteAddr().String())
		currentTime := uint64(time.Now().Unix())
		client := websocketserver.NewClient(conn.RemoteAddr().String(), conn, currentTime)
		go client.Read()
		go client.Write()
	}
}
