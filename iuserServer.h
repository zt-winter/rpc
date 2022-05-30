#include "user.h"

class IUserServer {
	public:
		virtual User findUserById(int id) = 0;
};

class UserServerImpl : public IUserServer {
	public:
		User findUserById(int id);
};
