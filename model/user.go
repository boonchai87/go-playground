package model

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" binding:"required" example:"choo@gmail.com" maxLength:"255"` //  E-mail
	Username string `json:"username" binding:"required" example:"choo" maxLength:"255"`        // Username
	//Password string `swaggerignore:"true" json:"password" binding:"required" example:"" maxLength:"255"`
}

type UserForCreate struct {
	User
	//ID int `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" binding:"required" example:"choo@gmail.com" maxLength:"255"` //  E-mail
	Username string `json:"username" binding:"required" example:"choo" maxLength:"255"`        // Username
	Password string `swaggerignore:"true" json:"password" binding:"required" example:"" maxLength:"255"`
}

type UserForUpdate struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" binding:"required" example:"choo@gmail.com" maxLength:"255"` //  E-mail
	Username string `json:"username" binding:"required" example:"choo" maxLength:"255"`        // Username
	//Password string `swaggerignore:"true" json:"password" binding:"required" example:"" maxLength:"255"`
}
