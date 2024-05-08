package password

import (
    "fmt"
    "strings"
    "reflect"
    "crypto/hmac"
    "crypto/sha512"
    "encoding/base64"
    "encoding/json"
//    "encoding/hex"
    "database/sql/driver"

    "gintemplate/utils/config"
)

type Password string

type password struct {
    Type string `json:"type"`
    Password string `json:"password"`
}

func New(pass string) Password {
    if pass == "" {
        return ""
    }
    salt := make([]byte, 16)
    _, err := rand.Read(salt)
    if err != nil {
        panic("Generate password error!")
    }
    h := hmac.New(sha512.New, []byte(config.Secret))
    //h := sha512.New()
    h.Write([]byte(pass))
    h.Write(salt)
    h.Sum(nil)
    return Password(fmt.Sprintf("$6$%s$%s", base64.StdEncoding.EncodeToString(salt), base64.StdEncoding.EncodeToString(hashedPassword)))
}

func (c Password) MarshalJSON() ([]byte, error) {
    pass := password{
        Type: "secret",
        Password: string(c),
    }
    return json.Marshal(pass)
}

func (c *Password) UnmarshalJSON(b []byte) error {
    var tmpstr string
    err := json.Unmarshal(b, &tmpstr)
    if err == nil {
        *c = New(tmpstr)
        return err
    }
    var tmppass password
    err = json.Unmarshal(b, &tmppass)
    if err != nil {
        return err
    }
    switch strings.ToLower(tmppass.Type) {
    case "plain":
        *c = New(tmppass.Password)
    case "secret":
        *c = Password(tmppass.Password)
    default:
        return fmt.Errorf("Invalid type %s", strings.ToLower(tmppass.Type))
    }
    return nil
}

func (c *Password) Scan(value interface{}) (err error) {
    if val, ok := value.(string); ok {
        *c = Password(val)
    } else {
        err = fmt.Errorf("sql: unsupported type %s", reflect.TypeOf(value))
    }
    return
}

func (c Password) Value() (driver.Value, error) {
    return string(c), nil
}
