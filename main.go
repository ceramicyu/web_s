package main

import (
	"io"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer,"Hi,this is a new world!")
	})
	http.ListenAndServe(":80",nil)
}
