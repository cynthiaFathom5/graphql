package database

var Names = []string{
	"James", "John", "Robert", "Michael", "William", "David", "Richard", "Joseph",
	"Thomas", "Charles", "Christopher", "Daniel", "Matthew", "Anthony", "Donald",
	"Mark", "Paul", "Steven", "Andrew", "Kenneth", "Joshua", "Kevin", "Brian", "George",
	"Edward", "Ronald", "Timothy", "Jason", "Jeffrey", "Ryan", "Jacob", "Gary", "Nicholas",
	"Eric", "Jonathan", "Stephen", "Larry", "Justin", "Scott", "Brandon", "Benjamin",
	"Samuel", "Frank", "Gregory", "Raymond", "Alexander", "Patrick", "Jack", "Dennis",
	"Jerry", "Tyler", "Aaron", "Jose", "Henry", "Adam", "Douglas", "Nathan", "Peter", "Zachary",
}

type User struct {
	isNode
	Name string
	ID   string
	Pets []Pet
}

var Users = make(map[string]User)

func AllUsersList() []*User {
	reply := make([]*User, len(Users))
	var i int
	for _, u := range Users {
		user := u
		reply[i] = &user
		i++
	}

	return reply
}
