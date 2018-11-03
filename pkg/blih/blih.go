// Go project by arthur
// blihUI
// 2018

package blih

import (
	"blihUI/pkg/data"
	"blihUI/pkg/user"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type BLIH struct {
	user           *user.User
	url, userAgent string
	verbose        bool
}

func New(url string, u *user.User, verbose bool, userAgent string) BLIH {
	return BLIH{user: u, url: url, userAgent: userAgent, verbose: verbose}
}

func (b *BLIH) Request(resource, method string, d *data.Data) (map[string]interface{}, error) {
	signed, err := data.Sign(b.user, d)
	if err != nil {
		return nil, err
	}
	marshaled, err := json.Marshal(signed)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, b.url+resource, bytes.NewBuffer(marshaled))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", b.userAgent)
	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var responseData map[string]interface{}
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("response code is not 200 OK but %s (%s)", resp.Status, responseData["error"]))
		return nil, err
	}
	return responseData, nil
}
