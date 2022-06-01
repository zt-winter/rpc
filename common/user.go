package common

import "strconv"

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

func NewUser(id int, name string) *User {
	return &User{Id: id, Name: name}
}

func (one *User) GetId() int {
	return one.Id
}

func (one *User) SetId(id int) {
	one.Id = id
}

func (one *User) GetName() string {
	return one.Name
}

func (one *User) SetName(name string) {
	one.Name = name
}

func (one *User) ToString() string {
	return "User{id=" + strconv.Itoa(one.Id) + ", name=" + one.Name + "}"
}

