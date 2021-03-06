// Go project by arthur
// glih
// 2018

package sshkey

import (
	"fmt"
	"glih/pkg/blih"
	"glih/pkg/data"
	"io/ioutil"
	"strings"
)

type SSHKey struct {
	name, key string
}

func Delete(name string, b *blih.BLIH) error {
	answer, err := b.Request("sshkey/"+name, "DELETE", nil)
	if err != nil {
		return err
	}
	fmt.Println(answer["message"].(string))
	return nil
}

func List(b *blih.BLIH) error {
	list, err := b.Request("sshkeys", "GET", nil)
	if err != nil {
		return err
	}
	for comment, key := range list {
		fmt.Printf("%s %s\n", key.(string), comment)
	}
	return nil
}

func Upload(filename string, b *blih.BLIH) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	stringFile := strings.TrimSuffix(string(file), "\n")
	key := escape(stringFile)
	d := &data.Data{"sshkey": key}
	answer, err := b.Request("sshkeys", "POST", d)
	if err != nil {
		return err
	}
	fmt.Println(answer["message"].(string))
	return nil
}

/*
	Functions below were taken straight from the golang net/url package and simplified for my needs.
	They are originally much more complex and versatile; but they gave different results from the python
	function quote() from the urllib.parse package, which is the kind of PathUrl quoting I am going for here.

	It was not my choice to disfigure such beautiful code. I did what had to be done to be compliant with
	the blih API. I am not proud of any of this and I still weep at nights.
 */

func shouldEscape(c byte) bool {
	if 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z' || '0' <= c && c <= '9' {
		return false
	}

	switch c {
	case '-', '_', '.', '/':
		return false
	case ';', '?', ':', '@', '&', '=', '+', '$', ',', '~':
		return true
	}

	return true
}

func escape(s string) string {
	hexCount := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if shouldEscape(c) {
			hexCount++
		}
	}

	if hexCount == 0 {
		return s
	}

	var buf [64]byte
	var t []byte

	required := len(s) + 2*hexCount
	if required <= len(buf) {
		t = buf[:required]
	} else {
		t = make([]byte, required)
	}

	if hexCount == 0 {
		copy(t, s)
		for i := 0; i < len(s); i++ {
			if s[i] == ' ' {
				t[i] = '+'
			}
		}
		return string(t)
	}

	j := 0
	for i := 0; i < len(s); i++ {
		switch c := s[i]; {
		case shouldEscape(c):
			t[j] = '%'
			t[j+1] = "0123456789ABCDEF"[c>>4]
			t[j+2] = "0123456789ABCDEF"[c&15]
			j += 3
		default:
			t[j] = s[i]
			j++
		}
	}
	return string(t)
}
