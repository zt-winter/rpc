#include "iuserServer.h"
#include <iostream>
#include <cstdlib>
#include <cstdio>
#include <cstring>
#include <sys/socket.h>
#include <netinet/in.h>
#include <sys/types.h>
#include <arpa/inet.h>

void process(int listenfd, int conn, bool* flag);

int main() {
	bool flag = true;
	int listenfd = 	socket(AF_INET, SOCK_STREAM, 0);
	if (listenfd == -1) {
		std::cout << "Error: socket" << std::endl;
	}
	struct sockaddr_in addr;
	addr.sin_family = AF_INET;
	addr.sin_port = htons(7020);
	addr.sin_addr.s_addr = inet_addr("127.0.0.1");
	if (bind(listenfd, (struct sockaddr*)&addr, sizeof(addr)) == -1) {
		std::cout << "Error: bind" << std::endl;
		return 0;
	}
	if (listen(listenfd, 5) == -1) {
		std::cout << "Error: listen" << std::endl;
		return 0;
	}
	struct sockaddr_in clientAddr;
	socklen_t clientAddrLen = sizeof(clientAddr);
	while(flag) {
		int conn = accept(listenfd, (struct sockaddr*)&clientAddr, &clientAddrLen);
		process(listenfd, conn, &flag);
	}
	return 0;
}

void process(int listenfd, int conn, bool* flag) {
	char bufRecv[5];
	char buf[255];
	memset(bufRecv, 0, sizeof(bufRecv));
	int len = recv(conn, bufRecv, sizeof(bufRecv), 0);
	if (len <= 0) {
		*flag = false;
		std::cout << "Error: recv" << std::endl;
		return ;
	}
	bufRecv[len] = '\0';
	int id = atoi(bufRecv);
	IUserServer* one = new UserServerImpl();
	one->findUserById(id);
}
