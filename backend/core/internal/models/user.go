package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID       uint32
	Username string
	Password string
	Email    string
}

func (u *User) GetID() uint32 {
	return u.ID
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetUserName() string {
	return u.Username
}

func (u *User) GetPassWord() string {
	return u.Password
}

// Function to handle user password

func (u *User) HashPassword() (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.GetPassWord()), 14)
	return string(bytes), err
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.GetPassWord()), []byte(password))
	return err == nil
}
