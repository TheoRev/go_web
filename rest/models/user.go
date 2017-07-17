package models

import "errors"

// User model
type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = make(map[int]User)

func SetDefaultUser() {
	user := User{
		Id:       1,
		Username: "theo",
		Password: "ambu26",
	}
	users[user.Id] = user
}

func GetUser(userId int) (User, error) {
	if user, ok := users[userId]; ok {
		return user, nil
	}
	return User{}, errors.New("El usuario no sencuentra dentro del map")
}
