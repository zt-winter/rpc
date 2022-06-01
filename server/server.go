package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"common"
)

func findUserById(id int) common.User {
	return *common.NewUser(id, "first")
}

func process(w http.ResponseWriter, req *http.Request) {
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

func main() {
	http.HandleFunc("/", process)
	http.ListenAndServe("127.0.0.1:7020", nil)
}

