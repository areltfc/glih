// Go project by arthur
// glih
// 2018

package user

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"syscall"
)

const prompt = "Password: "

type User struct {
	email, token string
}

func New(email, givenToken string) *User {
	return &User{email: email, token: givenToken}
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Token() string {
	return u.token
}

func (u *User) CalculateToken() {
	fmt.Fprintf(os.Stderr, prompt)
	password, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(os.Stderr)
	hasher := sha512.New()
	hasher.Write(password)
	u.token = hex.EncodeToString(hasher.Sum(nil))
}
