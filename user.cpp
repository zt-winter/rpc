#include <stdlib.h>
#include "user.h"

int User::getId(){
	return this->id;
}

void User::setId(int id){
	this->id = id;
}

string User::getNmae(){
	return this->name;
}

void User::setName(string name){
	this->name = name;
}

string User::toString(){
	return "User{id=" + to_string(id) + ", name=" + name + "}";
}

