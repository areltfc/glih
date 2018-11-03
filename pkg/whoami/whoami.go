// Go project by arthur
// glih
// 2018

package whoami

import (
	"fmt"
	"glih/pkg/blih"
)

type User string

func WhoAmI(b *blih.BLIH) error {
	identity, err := b.Request("whoami", "GET", nil)
	if err != nil {
		return err
	}
	fmt.Println(identity["message"].(string))
	return nil
}

func Execute(baseurl, user, token, userAgent string, verbose bool) error {
	b := blih.New(baseurl, userAgent, user, token, verbose)
	return WhoAmI(&b)
}