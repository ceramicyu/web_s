package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os/exec"
	"io"
)

func main(){
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer,"Hi,this is a new world! update 2  33")
	})
	http.HandleFunc("/git", func(writer http.ResponseWriter, request *http.Request) {
		cmd := exec.Command("/bin/bash", "-c", "")
		var out bytes.Buffer

		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			fmt.Println("git pull err :",err)
		}
		io.WriteString(writer,"git pull success")

	})
	http.ListenAndServe(":8080",nil)
}
