package types

type Student struct {
	id int
	Name string		`validate:"required"`
	Age int				`validate:"required"`
	Email string	`validate:"required"`
}