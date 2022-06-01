package main

import (
	"bytes"
	"common"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

func main() {
	var id int
	fmt.Println("输入用户id")
	fmt.Scanln(&id)
	userTransfer(id)
}

func userTransfer(id int) {
	conn, err := net.Dial("tcp", "127.0.0.1:7020")
	if err != nil {
		fmt.Println("dial error")
	}
	defer conn.Close()
	user := &common.User{
		Id: id,
		Name: "",
	}
	userJson, err := json.Marshal(user)
	_, err = conn.Write(userJson)
	if err != nil {
		fmt.Println("write error")
	}
	buf := make([]byte, 1024)
	num, err := conn.Read(buf)
	var recvUser common.User
	if err := json.Unmarshal(buf[:num], &recvUser); err != nil {
		fmt.Println("json Unmarshal error")
	}
	fmt.Println(recvUser.ToString())
}

func findUserById(id int) common.User {
	request := common.User{
		Id : id,
		Name : "",
	}
	requestBody := new(bytes.Buffer)
	json.NewEncoder(requestBody).Encode(request)
	url := "http://127.0.0.1:7020"
	req, err := http.NewRequest("POST", url, requestBody)
	if err != nil {
		fmt.Println("request error")
	}
	req.Header.Set("Content-Type", "application/json")
	one := &http.Client{}
	rsp, err := one.Do(req)
	if err != nil {
		fmt.Println("do request error")
	}
	defer rsp.Body.Close()
	recvBody, _ := ioutil.ReadAll(rsp.Body)
	var recvUser common.User
	if err := json.Unmarshal(recvBody, &recvUser); err != nil {
		fmt.Println("json error")
	}
	fmt.Println(recvUser.ToString())
	return recvUser
}
