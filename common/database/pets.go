package database

var Species = []string{
	"dog", "cat", "mouse", "horse",
}

type Pet struct {
	Name    string
	Species string
}

var Pets = make(map[string]Pet)

func AllPetsList() []*Pet {
	reply := make([]*Pet, len(Pets))
	var i int
	for _, p := range Pets {
		reply[i] = &p
		i++
	}

	return reply
}
