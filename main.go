package main

import (
	"fmt"

	"google.golang.org/grpc"
)

func main() {
	c := hello3(hello2(hello1()))
	fmt.Println(c("its c"))
	fmt.Println("commit5")

}

func ReturnError() {
	fmt.Println("error")
}

type Server struct {
	gRPCConn *grpc.ClientConn
}

func hello1() func(string) int {
	return (func(string2 string) int {
		fmt.Println(string2)
		return 10
	})
}

func hello2(next func(string) int) func(string) int {
	a := next("calling from hello2")
	return (func(string2 string) int {
		fmt.Println(string2)
		return 10 + a
	})
}

func hello3(next func(string) int) func(string) int {
	a := next("calling from hello3")
	return (func(s string) int {
		fmt.Println(s)
		return 10 + a

	})
}
