// Go project by arthur
// glih
// 2018

package sshkey

import (
	"glih/pkg/blih"
	"glih/pkg/data"
	"fmt"
	"io/ioutil"
	"strings"
)

type SSHKey struct {
	name, key string
}

func (s SSHKey) String() string {
	return fmt.Sprintf("%s %s", s.key, s.name)
}

func Delete(name string, b *blih.BLIH) error {
	_, err := b.Request("sshkey/"+name, "DELETE", nil)
	return err
}

func List(b *blih.BLIH) ([]SSHKey, error) {
	list, err := b.Request("sshkeys", "GET", nil)
	if err != nil {
		return nil, err
	}
	var keys []SSHKey
	for key, value := range list {
		keys = append(keys, SSHKey{name: key, key: value.(string)})
	}
	return keys, nil
}

func Upload(filename string, b *blih.BLIH) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	stringFile := strings.TrimSuffix(string(file), "\n")
	key := escape(stringFile)
	d := &data.Data{"sshkey": key}
	_, err = b.Request("sshkeys", "POST", d)
	return err
}

/*
	Functions below were take straight from the golang net/url package and simplified for my needs.
	They are originally much more complex and versatile; but they gave different results from the python
	function quote() from the urllib.parse package, which is the kind of PathUrl quoting I am going for here.

	It was not my choice to disfigure such beautiful code. I did what had to be done to be compliant with
	the blih API. I am not proud of any of this and I still weep about it at nights.
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
