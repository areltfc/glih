// Go project by arthur
// blihUI
// 2018

package token

import (
	"crypto/sha512"
	"encoding/hex"
)

type Token string

func (t *Token) ToSha512() string {
	hasher := sha512.New()
	hasher.Write([]byte(*t))
	return hex.EncodeToString(hasher.Sum(nil))
}
