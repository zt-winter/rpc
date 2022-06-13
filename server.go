package rpc

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"reflect"
	"sync"
)

func NewServer() *Server{
	return &Server{}
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
	this.serverMap.LoadOrStore(s.name, s)
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


func (this *Server) process(conn net.Conn) {
	defer conn.Close()
	send := new(sync.Mutex)
	wg := new(sync.WaitGroup)
	this.reqLock.Lock()
	this.req = new(Request)
	this.reqLock.Unlock()

	this.readRequest(conn, this.req)
	s1, _ := this.serverMap.Load(this.req.MethodName)
	s, ok := s1.(*service)
	if !ok {
		fmt.Println("the service is not exist")
		return 
	}
	meth := s.methodMap[this.req.MethodName]
	// inputParams may be not one, so if inputParams are not noe, ArgT is Pointer
	var argv reflect.Value
	if meth.ArgT.Kind() == reflect.Pointer {
		argv = reflect.New(meth.ArgT.Elem())
	} else {
		argv = reflect.New(meth.ArgT)
	}
	retv := reflect.New(meth.RetT.Elem())

	s.call(this, send, wg, meth, argv, retv)

	this.rspLock.Lock()
	this.rsp = new(Response)
	this.reqLock.Unlock()

	this.sendResponse(conn, this.rsp, retv)
	
	wg.Wait()
	conn.Close()
}

func (this *service) call(s *Server, send *sync.Mutex, wg *sync.WaitGroup, meth *methodType, argv reflect.Value, retv reflect.Value) {
	if wg != nil {
		defer wg.Done()
	}
	meth.Lock()
	meth.numCalls++
	meth.Unlock()
	function := meth.method.Func
	// reflect.Method.Func {params[0]: struct self, params[1]:input params, params[2]:output params}
	function.Call([]reflect.Value{this.recvV, argv, retv})
}

func (this *Server) readRequest(conn net.Conn, req *Request) {
	messageJson := make([]byte, 1024)
	num, err := conn.Read(messageJson)
	if err != nil {
		return 
	}
	err = json.Unmarshal(messageJson[:num], req)
	if err != nil {
		fmt.Println("json unmarshal error")
	}
}

func (this *Server) sendResponse(conn net.Conn, rsp *Response, retv interface{}) {
	rspJson, _ := json.Marshal(rsp)
	_, err := conn.Write(rspJson)
	if err != nil {
		fmt.Println("ret send error")
	}
}
