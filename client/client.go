package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"common"
)

func main() {
	findUserById(123)
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
