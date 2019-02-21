package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main(){
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer,"Hi,this is a new world! update 2019")
	})
	http.HandleFunc("/git", func(writer http.ResponseWriter, request *http.Request) {
		shellPath := "/home/www/web/web_s/r.sh"
		argv := make([]string, 1)
		attr := new(os.ProcAttr)
		newProcess, err := os.StartProcess(shellPath, argv, attr)  //运行脚本
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Process PID", newProcess.Pid)
		processState, err := newProcess.Wait() //等待命令执行完
		if err != nil {
			fmt.Println(err)
			io.WriteString(writer,"git pull err :"+fmt.Sprintf("%v",err))
		}else{
			io.WriteString(writer,"git pull success")
		}
		fmt.Println("processState PID:", processState.Pid())//获取PID
		fmt.Println("ProcessExit:", processState.Exited())//获取进程是否退出
		
	})
	http.ListenAndServe(":8080",nil)
}
