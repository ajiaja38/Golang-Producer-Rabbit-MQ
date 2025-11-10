package model

type UserDao struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Age         int64  `json:"age"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
}
