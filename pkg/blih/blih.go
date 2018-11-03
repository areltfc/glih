// Go project by arthur
// glih
// 2018

package blih

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"glih/pkg/data"
	"glih/pkg/user"
	"io/ioutil"
	"net/http"
)

type BLIH struct {
	user           *user.User
	url, userAgent string
	verbose        bool
}

func New(url, userAgent, username, token string, verbose bool) BLIH {
	u := user.New(username, token)
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
		err = errors.New("Can't decode data, aborting")
		return nil, err
	}
	if resp.StatusCode != 200 {
		errString := fmt.Sprintf("HTTP Error %d\n", resp.StatusCode)
		errString += fmt.Sprintf("Error message : '%s'\n", responseData["error"])
		err = errors.New(errString)
		return nil, err
	}
	return responseData, nil
}
