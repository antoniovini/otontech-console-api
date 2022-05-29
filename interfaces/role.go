package interfaces

type RoleManagment struct {
	Username string
	Role     string
}

type Role struct {
	Name        string
	Description string
	Level       uint
}
