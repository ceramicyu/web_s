FROM golang:latest

WORKDIR  $GOPATH/src/
ADD . $GOPATH/src/


RUN go build main.go

EXPOSE 8080

ENTRYPOINT ["./main"]