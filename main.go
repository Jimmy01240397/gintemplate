package main
import (
//    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/sessions"
    "github.com/gin-contrib/sessions/cookie"
    "github.com/go-errors/errors"

    "gintemplate/router"
    "gintemplate/utils/config"
    "gintemplate/utils/database"
    "gintemplate/utils/errutil"
    "gintemplate/middlewares/auth"
)

func main() {
    defer database.Close()
    store := cookie.NewStore([]byte(config.Secret))
    backend := gin.Default()
    backend.Use(errorHandler)
    backend.Use(gin.CustomRecovery(panicHandler))
    backend.Use(sessions.Sessions(config.Sessionname, store))
    backend.Use(auth.AddMeta)
    router.Init(&backend.RouterGroup)
    backend.Run(":"+string(config.Port))
}

func panicHandler(c *gin.Context, err any) {
    goErr := errors.Wrap(err, 2)
    errutil.AbortAndError(c, &errutil.Err{
        Code: 500,
        Msg: "Internal server error",
        Data: goErr.Error(),
    })
}

func errorHandler(c *gin.Context) {
    c.Next()

    for _, e := range c.Errors {
        err := e.Err
        if myErr, ok := err.(*errutil.Err); ok {
            if myErr.Msg != nil {
                c.JSON(myErr.Code, myErr.ToH())
            } else {
                c.Status(myErr.Code)
            }
        } else {
            c.JSON(500, gin.H{
                "code": 500,
                "msg": "Internal server error",
                //"data": err.Error(),
            })
        }
        return
    }
}
