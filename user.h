#include <string>
using namespace std;

class User {
	private:
		int id;
		string name;
	public:
		int getId();
		void setId(int id);
		string getNmae();
		void setName(string name);
		string toString();
	public:
		User(int id, string name) {
			this->id = id;
			this->name = name;
		}
};
