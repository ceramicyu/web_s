package main

import (
	"net/http"
	"github.com/gorilla/websocket"
	"imp"
	"fmt"
)
var  Poll []*imp.Wsmsg

func wsHandler(w http.ResponseWriter , r *http.Request){
	fmt.Println(r.RemoteAddr)
	up:=&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	var (
		wsconn *websocket.Conn
		err error

	)

	if wsconn,err= up.Upgrade(w,r,nil);err!=nil{
		return
	}
	ws:=imp.Initwsmsg(wsconn)
	Poll=append(Poll, ws)

	go ws.WriteToAll(&Poll)
	go ws.Read()

	return
}

func main(){
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe("10.0.3.177:5555",nil)
}
