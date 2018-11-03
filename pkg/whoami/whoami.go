// Go project by arthur
// glih
// 2018

package whoami

import (
	"fmt"
	"glih/pkg/blih"
	"os"
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

func Execute(args []string, baseurl, user, token, userAgent string, verbose bool) error {
	if len(args) != 0 {
		fmt.Fprintf(os.Stderr, "Too many arguments for command 'whoami'\n")
		os.Exit(1)
	}
	b := blih.New(baseurl, userAgent, user, token, verbose)
	return WhoAmI(&b)
}