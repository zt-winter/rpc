package rpc

import (
	"reflect"
	"sync"
)

type Request struct {
	ServiceName string `json:"servicename"`
	MethodName string `json:"methodname"`
	Args	interface{} `json:"args"`
}

type methodType struct {
	sync.Mutex
	numCalls int
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
	//serverMap map[string]*service
	serverMap sync.Map
	reqLock sync.Mutex
	req *Request
	rspLock sync.Mutex
	rsp *Response
}


type Response struct {
	Ret interface{} `json:"ret"`
}

