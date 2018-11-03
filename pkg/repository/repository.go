// Go project by arthur
// glih
// 2018

package repository

import (
	"encoding/json"
	"fmt"
	"glih/pkg/blih"
	"glih/pkg/data"
	"os"
	"time"
)

type Repository struct {
	name, uuid, url, description string
	public                       bool
	creation                     time.Time
	acl                          map[string]string
}

func List(b *blih.BLIH) error {
	repositories, err := b.Request("repositories", "GET", nil)
	if err != nil {
		return err
	}
	list := repositories["repositories"].(map[string]interface{})
	for repository := range list {
		fmt.Println(repository)
	}
	return nil
}

func Create(name, description string, b *blih.BLIH) error {
	d := data.Data{"name": name, "type": "git"}
	if description != "" {
		d["description"] = description
	}
	answer, err := b.Request("repositories", "POST", &d)
	fmt.Println(answer["message"].(string))
	return err
}

func Delete(name string, b *blih.BLIH) error {
	answer, err := b.Request("repository/"+name, "DELETE", nil)
	fmt.Println(answer["message"].(string))
	return err
}

func Info(name string, b *blih.BLIH) error {
	repository, err := b.Request("repository/"+name, "GET", nil)
	if err != nil {
		return err
	}
	marshaled, err := json.Marshal(repository["message"])
	if err != nil {
		return err
	}
	fmt.Println(string(marshaled))
	return nil
}

func SetACL(name, acluser, acl string, b *blih.BLIH) error {
	d := data.Data{"user": acluser, "acl": acl}
	answer, err := b.Request("repository/"+name+"/acls", "POST", &d)
	fmt.Println(answer["message"].(string))
	return err
}

func GetACL(name string, b *blih.BLIH) error {
	repository, err := b.Request("repository/"+name+"/acls", "GET", nil)
	if err != nil {
		return err
	}
	for user, acls := range repository {
		fmt.Printf("%s:%s", user, acls.(string))
	}
	return nil
}

func repositoryUsage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [options] repository command ...\n\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Commands:\n")
	fmt.Fprintf(os.Stderr, "\tcreate repo\t\t\t-- Create a repository named \"repo\"\n")
	fmt.Fprintf(os.Stderr, "\tinfo repo\t\t\t-- Get the repository metadata\n")
	fmt.Fprintf(os.Stderr, "\tlist\t\t\t\t-- List the repositories created\n")
	fmt.Fprintf(os.Stderr, "\tsetacl repo user [acl]\t\t-- Set (or remove) an acl for \"user\" on \"repo\"\n")
	fmt.Fprintf(os.Stderr, "\t\t\t\t\tACL format:\n")
	fmt.Fprintf(os.Stderr, "\t\t\t\t\t\tr for read\n")
	fmt.Fprintf(os.Stderr, "\t\t\t\t\t\tw for write\n")
	fmt.Fprintf(os.Stderr, "\t\t\t\t\t\ta for admin\n")
	fmt.Fprintf(os.Stderr, "\tgetacl repo\t\t\t-- Get the acls set for the repository\n")
	os.Exit(1)
}

func Execute(args []string, baseurl, user, token, userAgent string, verbose bool) error {
	argsLen := len(args)
	var err error
	if argsLen == 0 {
		repositoryUsage()
	}
	switch args[0] {
	case "create":
		if argsLen != 2 {
			repositoryUsage()
		}
		b := blih.New(baseurl, userAgent, user, token, verbose)
		err = Create(args[1], "", &b)
	case "list":
		if argsLen != 1 {
			repositoryUsage()
		}
		b := blih.New(baseurl, userAgent, user, token, verbose)
		err = List(&b)
	case "info":
		if argsLen != 2 {
			repositoryUsage()
		}
		b := blih.New(baseurl, userAgent, user, token, verbose)
		err = Info(args[1], &b)
	case "delete":
		if argsLen != 2 {
			repositoryUsage()
		}
		b := blih.New(baseurl, userAgent, user, token, verbose)
		err = Delete(args[1], &b)
	case "setacl":
		var acl string
		if argsLen == 3 {
			acl = ""
		} else if argsLen == 4 {
			acl = args[3]
		} else {
			repositoryUsage()
		}
		b := blih.New(baseurl, userAgent, user, token, verbose)
		err = SetACL(args[1], args[2], acl, &b)
	case "getacl":
		if argsLen != 2 {
			repositoryUsage()
		}
		b := blih.New(baseurl, userAgent, user, token, verbose)
		err = GetACL(args[1], &b)
	default:
		repositoryUsage()
	}
	return err
}