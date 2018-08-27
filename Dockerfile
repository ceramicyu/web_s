FROM golang:onbuild
EXPOSE 5555
RUN go get github.com/gorilla/websocket
RUN export GOPATH=~/web_s
