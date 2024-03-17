package auth

import (
//    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/sessions"

    "gintemplate/utils/errutil"
    "gintemplate/models/user"
)

func CheckSignIn(c *gin.Context) {
    if isSignIn, exist := c.Get("isSignIn"); !exist || !isSignIn.(bool) {
        errutil.AbortAndStatus(c, 401)
    }
}

func CheckIsAdmin(c *gin.Context) {
    if isAdmin, exist := c.Get("isAdmin"); !exist || !isAdmin.(bool) {
        errutil.AbortAndStatus(c, 401)
    }
}

func AddMeta(c *gin.Context) {
    session := sessions.Default(c)
    username := session.Get("user")
    if username == nil {
        c.Set("isSignIn", false)
    } else {
        userdata, _ := user.GetUser(username.(string))
        if userdata == nil {
            c.Set("isSignIn", false)
        } else {
            c.Set("isSignIn", true)
            c.Set("isAdmin", userdata.IsAdmin())
        }
    }
}
