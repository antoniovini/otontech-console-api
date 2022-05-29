package interfaces

type CreateStep struct {
	Name   string
	Params []Param
}

type UpdateStep struct {
	Name   string
	Params []Param
}

type Step struct {
	Name   string
	Params []Param
}
