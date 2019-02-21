package main

import (
	"fmt"
	"io"
	"net/http"
	"os/exec"
)

func main(){
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer,"Hi,this is a new world! update 2019 2 ")
	})
	http.HandleFunc("/git", func(writer http.ResponseWriter, request *http.Request) {
		shellPath := "$GOPATH/src/r.sh"
		command := exec.Command(shellPath) //初始化Cmd
		err := command.Start()//运行脚本		
		if err != nil {
			fmt.Println(err)
			io.WriteString(writer,"git pull err :"+fmt.Sprintf("%v",err))
		}else{
			io.WriteString(writer,"git pull success")
		}
		

	})
	http.ListenAndServe(":8080",nil)
}

