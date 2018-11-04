// Go project by arthur
// glih
// 2018

package data

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"glih/pkg/user"
)

type Data map[string]interface{}

func sign(u *user.User, data *Data) (Data, error) {
	mac := hmac.New(sha512.New, []byte(u.Token()))
	update := []byte(u.Email())
	if data != nil {
		dump, err := json.MarshalIndent(*data, "", "    ")
		if err != nil {
			signed := make(Data)
			return signed, err
		}
		update = append(update, dump...)
	}
	mac.Write(update)
	signed := Data{"user": u.Email(), "signature": hex.EncodeToString(mac.Sum(nil))}
	if data != nil {
		signed["data"] = data
	}
	return signed, nil
}

func TreatForHTTPRequest(u *user.User, d *Data) ([]byte, error) {
	if u.Token() == "" {
		u.CalculateToken()
	}
	signed, err := sign(u, d)
	if err != nil {
		return nil, err
	}
	marshaled, err := json.Marshal(signed)
	if err != nil {
		return nil, err
	}
	return marshaled, nil
}
