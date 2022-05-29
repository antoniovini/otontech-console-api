package interfaces

type CreateArg struct {
	Name        string
	Required    bool
	Identifier  string
	Description string
}

type UpdateArg struct {
	Name        string
	Required    bool
	Identifier  string
	Description string
}

type Arg struct {
	Name        string
	Required    bool
	Identifier  string
	Description string
}
