#include "user.h"
#include <cstdio>
#include <stdlib.h>
#include <cstring>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <iostream>

int main() {
	int clientFd;
	struct sockaddr_in serverAddr;
	char buf[255];
	memset(buf, 0, 255);
	memset(&serverAddr, 0, sizeof(serverAddr));
	serverAddr.sin_family = AF_INET;
	serverAddr.sin_addr.s_addr = inet_addr("127.0.0.1");
	serverAddr.sin_port = htons(7020);
	if ((clientFd = socket(PF_INET, SOCK_STREAM, 0)) < 0) {
		std::cout << "socket error" << std::endl;
		return 0;
	}
	if (connect(clientFd, (struct sockaddr*)&serverAddr, sizeof(struct sockaddr)) < 0) {
		std::cout << "listen error" << std::endl;
		return 0;
	}
	printf("connected to server\n");
	sprintf(buf, "%d", 123);
	printf("%s\n", buf);
	printf("%d\n", (int)strlen(buf));
	send(clientFd, buf, strlen(buf), 0);
	int len = recv(clientFd, buf, 200, 0);
	if (len <= 0) {
		std::cout << "recv error" << std::endl;
		return 0;
	}
	buf[len] = '\0';
	printf("%s\n", buf);
	User one = User(123, string(buf));
	std::cout << one.toString() << std::endl;
	return 0;
}
