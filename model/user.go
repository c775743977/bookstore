package model

type User struct {
	ID int `form:"userID"`
	Name string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Email string `form:"email" binding:"required"`
	Privilege string `form:"privilege" binding:"required"`
	Orders []*Order
	Err string `gorm:"-"`
}

func (user *User) IsRoot() bool {
	if user.Privilege == "Y" {
		return true
	} else {
		return false
	}
}