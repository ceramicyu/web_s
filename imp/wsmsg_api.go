package imp

import (
	"github.com/gorilla/websocket"
	"fmt"
)

type Wsmsg struct {
	conn *websocket.Conn
	rc chan []byte
	wc chan []byte
}

func Initwsmsg(conn *websocket.Conn)(*Wsmsg){
	ws:=&Wsmsg{conn:conn,
	           rc:make(chan []byte),
	           wc:make(chan []byte),
	                }
	return ws
}
func (ws *Wsmsg)Read(){
	var (
		data []byte
		err error
	)
	for{
		if _,data,err = ws.conn.ReadMessage();err!=nil{
			return
		}
		fmt.Println("recieve msg >>  ",string(data))
		ws.wc<-[]byte("【"+ws.conn.RemoteAddr().String()+"】:"+string(data) )
	}

}
func (ws *Wsmsg)Write(){
	for{
		data:=<-ws.wc
		ws.conn.WriteMessage(websocket.TextMessage,data)
	}

}
func (ws *Wsmsg)WriteToAll(Poll *[]*Wsmsg){
	for{
		data:=<-ws.wc
		 for _,v:=range *Poll{
			v.conn.WriteMessage(websocket.TextMessage,data)
		 }

	}

}
func (ws *Wsmsg)close(){

}