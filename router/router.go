package router
import (
    "github.com/gin-gonic/gin"

    "gintemplate/middlewares/auth"
)

var router *gin.RouterGroup

func Init(r *gin.RouterGroup) {
    router = r
    router.GET("/status", auth.CheckSignIn, status)
    //user.Init(router.Group("/user"))
}

func status(c *gin.Context) {
    c.String(200, "test2")
}
