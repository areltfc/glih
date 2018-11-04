// +build windows

package sshkey

import (
	"fmt"
	"glih/pkg/blih"
	"os"
)

func sshkeyUsage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [options] sshkey command ...\n\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Commands:\n")
	fmt.Fprintf(os.Stderr, "\tupload <file>\t\t\t-- Upload a new ssh-key located in the file <file>\n")
	fmt.Fprintf(os.Stderr, "\tlist\t\t\t\t-- List the ssh-keys\n")
	fmt.Fprintf(os.Stderr, "\tdelete <sshkey>\t\t\t-- Delete the sshkey with comment <sshkey>\n")
	os.Exit(1)
}

func Execute(args []string, baseURL, user, token, userAgent string) error {
	argsLen := len(args)
	var err error
	if argsLen == 0 {
		sshkeyUsage()
	}
	b := blih.New(baseURL, userAgent, user, token)
	switch args[0] {
	case "list":
		if argsLen > 1 {
			sshkeyUsage()
		}
		err = List(b)
	case "upload":
		if argsLen != 2 {
			sshkeyUsage()
		}
		err = Upload(args[1], b)
	case "delete":
		if argsLen != 2 {
			sshkeyUsage()
		}
		err = Delete(args[1], b)
	}
	return err
}
