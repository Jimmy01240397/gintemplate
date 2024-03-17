package password

import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/base64"
//    "encoding/hex"

    "github.com/go-errors/errors"

    "gintemplate/utils/config"
)

type Password struct {
    encryptpass string
}

func New(password string) Password {
    h := hmac.New(sha256.New, []byte(config.Secret))
    h.Write([]byte(password))
    return Password{
        encryptpass: base64.StdEncoding.EncodeToString(h.Sum(nil)),
    }
}

func NewEncrypt(encryptpass string) Password {
    return Password{
        encryptpass: string(encryptpass),
    }
}

func (pass Password) String() string {
    return pass.encryptpass
}

func (pass *Password) Scan(src any) (err error) {
    var ok bool
    pass.encryptpass, ok = src.(string)
    if !ok {
        err = errors.New("Bad password type")
    }
    return
}
