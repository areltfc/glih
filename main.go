// Go project by arthur
// glih
// 2018

package main

import (
	"glih/pkg/blih"
	"glih/pkg/user"
)

const (
	version       = "1.7"
	email         = ""
	baseURL       = "https://blih.epitech.eu/"
	baseUserAgent = "blih-" + version
)

func main() {
	u := user.New(email)
	_ = blih.New(baseURL, u, false, baseUserAgent)
}
