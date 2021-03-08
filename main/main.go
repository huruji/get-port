package main

import (
"fmt"
"net"
	"reflect"
)


func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "localhost:0")
	l, _ := net.ListenTCP("tcp", addr)
	fmt.Println(l.Addr().(*net.TCPAddr).Port)
	fmt.Println(reflect.TypeOf(l.Addr()))
}

