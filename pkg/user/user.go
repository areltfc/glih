// Go project by arthur
// glih
// 2018

package user

import (
	"fmt"
	"glih/pkg/token"
	"golang.org/x/crypto/ssh/terminal"
	"syscall"
)

const prompt = "Password: "

type User struct {
	email, token string
}

func New(email, givenToken string) *User {
	fmt.Printf(prompt)
	password, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		panic(err)
	}
	fmt.Println()
	var t string
	if givenToken == "" {
		tok := token.Token(password)
		t = tok.ToSha512()
	} else {
		t = givenToken
	}
	return &User{email: email, token: t}
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Token() string {
	return u.token
}
