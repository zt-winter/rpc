package rpc

import (
	"net"
	"reflect"
)

type Request struct {
	ServiceName string `json:"servicename"`
	MethodName string `json:"methodname"`
	Args	interface{} `json:"args"`
}

type methodType struct {
	method reflect.Method
	ArgT reflect.Type
	RetT reflect.Type
}

type service struct {
	name string
	recvV reflect.Value
	recvT reflect.Type
	methodMap map[string]*methodType
}

type Server struct {
	l net.Listener
	serverMap map[string]*service
}


type Response struct {
	Ret interface{} `json:"ret"`
}

