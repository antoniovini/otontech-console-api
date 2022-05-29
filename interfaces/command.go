package interfaces

type CreateCommand struct {
	Description string
	Activator   string
	Action      string
	Roles       []string
	Steps       []Step
	Args        []Arg
}

type UpdateCommand struct {
	Description string
	Activator   string
	Action      string
	Roles       []string
	Steps       []Step
	Args        []Arg
}

type Command struct {
	Id          uint
	UniqueId    string
	Description string
	Activator   string
	Action      string
	Roles       []Role
	Steps       []Step
	Args        []Arg
}
