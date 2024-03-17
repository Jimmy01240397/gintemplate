package tablename

import (
    "log"
    "fmt"
    "time"
    "strings"
    "encoding/json"

    "github.com/go-errors/errors"

    "gintemplate/utils/database"
    "gintemplate/utils/password"
)

type User struct {
    Username string
    Password password.Password
    Online bool
    Enable bool
    Startdate *time.Time
    Enddate *time.Time
}

func (user *User) ToMap() map[string]any {
    var usermap map[string]any
    userjson, _ := json.Marshal(user)
    json.Unmarshal(userjson, &usermap)
    usermap["Password"] = user.Password.String()
    return usermap
}

const tablename = "user"

func init() {
    _, err := database.Exec(fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %s (
            username varchar(255) NOT NULL PRIMARY KEY,
            password varchar(255),
            online tinyint(1) NOT NULL DEFAULT 0,
            enable tinyint(1) NOT NULL DEFAULT 1,
            startdate date DEFAULT NULL,
            enddate date DEFAULT NULL
        )
    `, tablename))
    if err != nil {
        log.Panicln(err)
        return
    }
}

