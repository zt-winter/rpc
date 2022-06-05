package rpc

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"reflect"
)

func NewServer() *Server{
	return &Server{}
}

func (this *Server) start() {
	listen, err := net.Listen("tcp", "127.0.0.1:7020")
	this.l = listen
	if err != nil {
		fmt.Println("listen error")
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept error")
		}
		go this.process(conn)
	}
}

func (this *Server) register(recv interface{}) {
	s := new(service)
	s.recvT = reflect.TypeOf(recv)
	s.recvV = reflect.ValueOf(recv)
	sname := reflect.Indirect(s.recvV).Type().Name()
	if sname == "" {
		log.Println("rpc register no export server name")
	}
	s.name = sname
	s.methodMap = methodMapBuild(s.recvT)
	this.serverMap[s.name] = s
}

func methodMapBuild(recvT reflect.Type) map[string]*methodType {
	methodMap := make(map[string]*methodType, 0)
	for i := 0; i < recvT.NumMethod(); i++ {
		method := recvT.Method(i)
		if !method.IsExported() {
			continue
		}
		mType := method.Type
		mName := method.Name
		argT := mType.In(1)
		retT := mType.In(2)
		if mType.NumOut() != 1 {
			continue
		}
		methodMap[mName] = &methodType{method:method, ArgT:argT, RetT:retT}
	}
	return methodMap
}


func (this *Server) stop() {
	this.l.Close()
}


func (this *Server) process(conn net.Conn) {
	defer conn.Close()
	messageJson := make([]byte, 1024)
	num, err := conn.Read(messageJson)
	if err != nil {
		return
	}
	var req Request
	err = json.Unmarshal(messageJson[:num], &req)
	if err != nil {
		fmt.Println("json unmarshal error")
	}
	s := this.serverMap[req.ServiceName]
	meth := s.methodMap[req.MethodName]
	function := meth.method.Func
	var argv reflect.Value
	// inputParams may be not one, so if inputParams are not noe, ArgT is Pointer
	if meth.ArgT.Kind() == reflect.Pointer {
		argv = reflect.New(meth.ArgT.Elem())
	} else {
		argv = reflect.New(meth.ArgT)
	}
	retv := reflect.New(meth.RetT.Elem())
	// reflect.Method.Func {params[0]: struct self, params[1]:input params, params[2]:output params}
	function.Call([]reflect.Value{s.recvV, argv, retv})
	rsp := &Response{
		Ret: retv.Interface(),
	}
	rspJson, _ := json.Marshal(rsp)
	_, err = conn.Write(rspJson)
	if err != nil {
		fmt.Println("ret send error")
	}
}
