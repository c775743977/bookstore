package model

type User struct {
	ID int
	Name string
	Password string
	Email string
	Privilege string
	Orders []*Order
}

func (user *User) IsRoot() bool {
	if user.Privilege == "Y" {
		return true
	} else {
		return false
	}
}