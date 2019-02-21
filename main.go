package main

import (
	"fmt"
	"io"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer,"<h1>Hi,this is a new world! update 2019 asdgajhgdjah</h1>")
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
