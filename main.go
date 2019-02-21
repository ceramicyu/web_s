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
		io.WriteString(writer,"Hi,this is a new world! update 2019")
	})
	http.HandleFunc("/git", func(writer http.ResponseWriter, request *http.Request) {
		cmd := exec.Command("/bin/bash", "-c", "./home/www/web/web_s/r.sh")
		var out bytes.Buffer

		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			fmt.Println("git pull err :",err)
			io.WriteString(writer,"git pull success"+fmt.Sprintf("%v",err))
		}else{
			io.WriteString(writer,"git pull success")
		}

	})
	http.ListenAndServe(":8080",nil)
}
