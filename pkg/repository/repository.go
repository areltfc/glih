// Go project by arthur
// glih
// 2018

package repository

import (
	"encoding/json"
	"fmt"
	"glih/pkg/blih"
	"glih/pkg/data"
	"time"
)

type Repository struct {
	name, uuid, url, description string
	public                       bool
	creation                     time.Time
	acl                          map[string]string
}

func (r *Repository) Name() string {
	return r.name
}

func (r *Repository) UUID() string {
	return r.UUID()
}

func (r *Repository) URL() string {
	return r.url
}

func (r Repository) String() string {
	return fmt.Sprintf("%s", r.name)
}

func List(b *blih.BLIH) error {
	repositories, err := b.Request("repositories", "GET", nil)
	if err != nil {
		return err
	}
	for key := range repositories {
		fmt.Println(key)
	}
	return nil
}

func Create(name, description string, b *blih.BLIH) error {
	d := data.Data{"name": name, "type": "git"}
	if description != "" {
		d["description"] = description
	}
	_, err := b.Request("repositories", "POST", &d)
	return err
}

func Delete(name string, b *blih.BLIH) error {
	_, err := b.Request("repository/"+name, "DELETE", nil)
	return err
}

func Info(name string, b *blih.BLIH) error {
	repository, err := b.Request("repository/"+name, "GET", nil)
	if err != nil {
		return err
	}
	infos := repository["message"].(map[string]interface{})
	repo := map[string]string{
		"name":          name,
		"uuid":          infos["uuid"].(string),
		"description":   infos["description"].(string),
		"url":           infos["url"].(string),
		"public":        infos["public"].(string),
		"creation_time": infos["creation_time"].(string),
	}
	marshaled, err := json.Marshal(repo)
	if err != nil {
		return err
	}
	fmt.Println(string(marshaled))
	return nil
}

func SetACL(name, acluser, acl string, b *blih.BLIH) error {
	d := data.Data{"user": acluser, "acl": acl}
	_, err := b.Request("repository/"+name+"/acls", "POST", &d)
	return err
}

func GetACL(name string, b *blih.BLIH) error {
	repository, err := b.Request("repository/"+name+"/acls", "GET", nil)
	if err != nil {
		return err
	}
	for user, acls := range repository {
		fmt.Printf("%s %s", user, acls.(string))
	}
	return nil
}
