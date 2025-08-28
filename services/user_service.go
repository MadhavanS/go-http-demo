package services

import "go-http-demo/models"

var users []models.User
var nextID = 1

func ListUsers() []models.User {
	return users
}

func CreateUser(u models.User) {
	u.ID = nextID
	nextID++
	users = append(users, u)
}
