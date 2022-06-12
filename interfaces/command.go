package interfaces

type CreateCommand struct {
	Description string
	Activator   string
	Action      string
	Role        string
	Args        []Arg
}

type UpdateCommand struct {
	Description string
	Activator   string
	Action      string
	Role        string
	Args        []Arg
}

type Command struct {
	Id          uint
	UniqueId    string
	Description string
	Activator   string
	Action      string
	Role        Role
	Args        []Arg
}
