package service

import "fmt"

var allUsers []*User

type User struct {
	ID       int     `json:"id"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Name     string  `json:"name"`
	Age      float32 `json:"age"`
	IsAdmin  bool    `json:"isAdmin"`
}

// func GetAllUsers(limit, offset int) ([]*User, int, error) {
// 	//business logic her

// 	return allUsers, 100, nil
// }
// func NewUser(
// 	Username string,
// 	Password string,
// 	Name string,
// 	Age float32,
// 	IsAdmin bool) (*User, error) {
// 	//validation
// 	return nil, errors.New("")

// 	// return &User{
// 	// 	Username: Username,
// 	// 	Password: Password,
// 	// 	Name:     Name,
// 	// 	Age:      Age,
// 	// 	IsAdmin:  IsAdmin,
// 	// }
// }
// func UpdateUserByID()  {
// 		//getByiD
// 		Username
// 		Password
// 		Name
// 		Age
// 		IsAdmin
// 		// user.Update(Username
// 		Password
// 		Name
// 		Age
// 		IsAdmin
// 	)
// }
// func (u *User)Update()  {
// 	u.id = u.id
// 	u.Age=Age
// 	u.Age=Age
// 	u.Age=Age
// 	u.Age=Age
// }
func GetUserByID() {
	fmt.Println("GetUserByID service Called")
}
