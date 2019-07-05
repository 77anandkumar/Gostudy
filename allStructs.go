package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net/http"
)

type Server struct {
	gRPCConn *grpc.ClientConn
	Server *http.Server
}

func main(){
	fmt.Println("hello")
	err :=
}