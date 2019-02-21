package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main(){
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		f,_:=os.Open("./public/html/index.html")
		buf:=make([]byte,1024*1024)
		n,_:=f.Read(buf)
		io.WriteString(writer,string(buf[:n]))
	})
	http.HandleFunc("/git", func(writer http.ResponseWriter, request *http.Request) {

         _,err:=  http.Get("http://106.13.60.133:55555/")
		if err != nil {			
			io.WriteString(writer,"git pull err :"+fmt.Sprintf("%v",err))
		}else{
			io.WriteString(writer,"git pull success")
		}

	})
	http.ListenAndServe(":8080",nil)
}
