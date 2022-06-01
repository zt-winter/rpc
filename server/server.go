package main

import (
	"common"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
)

func findUserById(id int) common.User {
	switch id {
	case 1:
		return *common.NewUser(id, "first")
	case 2:
		return *common.NewUser(id, "second")
	default:
		return *common.NewUser(0, "zero")
	}
}

func processHttp(w http.ResponseWriter, req *http.Request) {
	recvBody, _ := ioutil.ReadAll(req.Body)
	var recvUser common.User
	if err := json.Unmarshal(recvBody, &recvUser); err != nil {
		fmt.Println("json error")
	}
	id := recvUser.GetId()
	user := findUserById(id)
	userJson, _ := json.Marshal(user)
	io.WriteString(w, string(userJson))
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		messageJson := make([]byte, 1024)
		num, err := conn.Read(messageJson)
		if err != nil {
			break
		}
		var recvUser common.User
		err = json.Unmarshal(messageJson[:num], &recvUser)
		if err != nil {
			fmt.Println("json unmarshal error")
		}
		id := recvUser.GetId()
		user := findUserById(id)
		uersJson, _ := json.Marshal(user)
		_, err = conn.Write(uersJson)
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:7020")
	if err != nil {
		fmt.Println("listen error")
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept error")
		}
		go process(conn)
	}
}

