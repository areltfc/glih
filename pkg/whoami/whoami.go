// Go project by arthur
// glih
// 2018

package whoami

import (
	"glih/pkg/blih"
)

type User string

func WhoAmI(b *blih.BLIH) (User, error) {
	identity, err := b.Request("whoami", "GET", nil)
	if err != nil {
		return "", err
	}
	user := identity["message"].(string)
	return User(user), nil
}
