#include "iuserServer.h"

User UserServerImpl::findUserById(int id) {
	return User(id, "sandy");
}
