package rpc

import (
	"encoding/json"
	"fmt"
	"net"
)

func Call(conn net.Conn, serviceName string, methodName string, args []interface{}) interface{} {
	req := &Request{
		ServiceName: serviceName,
		MethodName: methodName,
		Args: args,
	}
	reqJson, _ := json.Marshal(req)
	conn.Write(reqJson)
	messageJson := make([]byte, 1024)
	num, err := conn.Read(messageJson)
	if err != nil {
		fmt.Println("conn read error")
		return nil
	}
	var rsp Response
	err = json.Unmarshal(messageJson[:num], &rsp)
	return rsp
}
