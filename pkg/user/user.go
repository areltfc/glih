// Go project by arthur
// glih
// 2018

package user

import (
	"glih/pkg/token"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"syscall"
)

const prompt = "Mot de passe bocal : "

type User struct {
	email, token string
}

func New(email string) *User {
	fmt.Printf(prompt)
	password, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		panic(err)
	}
	fmt.Println()
	t := token.Token(password)
	return &User{email: email, token: t.ToSha512()}
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Token() string {
	return u.token
}
