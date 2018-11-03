// Go project by arthur
// glih
// 2018

package main

import (
	"fmt"
	"github.com/pborman/getopt/v2"
	"glih/pkg/repository"
	"glih/pkg/sshkey"
	"glih/pkg/whoami"
	"os"
)

const (
	version       = "1.7"
	baseURL       = "https://blih.epitech.eu/"
	baseUserAgent = "blih-" + version
)

var (
	v    = getopt.BoolLong("verbose", 'v', "")
	V    = getopt.BoolLong("version", 'V', "")
	u, t string
	U    = baseUserAgent
	b    = baseURL
)

func usage() {
	fmt.Printf("Usage: %s [options] command ...\n\n", os.Args[0])
	fmt.Printf("Global Options: \n")
	fmt.Printf("\t-u user\t| --user=user\t\t-- Run as user\n")
	fmt.Printf("\t-v\t| --verbose\t\t-- Verbose\n")
	fmt.Printf("\t-b url\t| --baseurl=url\t\t-- Base URL for BLIH\n")
	fmt.Printf("\t-t\t| --token\t\t-- Specify token in the cmdline\n")
	fmt.Printf("\t-V\t| --version\t\t-- Print this binary's version then exit\n")
	fmt.Printf("\t-U\t| --useragent=\t\t-- User-Agent for BLIH\n\n")
	fmt.Printf("Commands:\n")
	fmt.Printf("\trepository\t\t\t-- Repository management\n")
	fmt.Printf("\tsshkey\t\t\t\t-- SSH-KEYS management\n")
	fmt.Printf("\twhoami\t\t\t\t-- Print who you are\n")
}

func init() {
	getopt.FlagLong(&u, "user", 'u', "")
	getopt.FlagLong(&U, "useragent", 'U', "")
	getopt.FlagLong(&b, "baseurl", 'b', "")
	getopt.FlagLong(&t, "token", 't', "")
}

func main() {
	err := getopt.Getopt(nil)
	if err != nil {
		fmt.Println(err)
		usage()
		os.Exit(1)
	}
	if *V == true {
		fmt.Println("blih version " + version)
		return
	}
	args := getopt.Args()
	if len(args) == 0 {
		usage()
		os.Exit(1)
	}
	switch args[0] {
	case "repository":
		err = repository.Execute(args[1:], b, u, t, U, *v)
	case "sshkey":
		err = sshkey.Execute(args[1:], b, u, t, U, *v)
	case "whoami":
		err = whoami.Execute(b, u, t, U, *v)
	default:
		usage()
		os.Exit(1)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
	}
}
