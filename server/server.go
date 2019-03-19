package server

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/static"

    "autochess/controller"
)

// Init is initialize server
func Init() {
    r := router()
    r.Run()
}

func router() *gin.Engine {
    r := gin.Default()
    r.LoadHTMLGlob("client/dist/*.html")
    r.Use(static.Serve("/js", static.LocalFile("./client/dist/js", true)))
    r.Use(static.Serve("/css", static.LocalFile("./client/dist/css", true)))
    r.Use(static.Serve("/img", static.LocalFile("./client/dist/img", true)))
    staticCtrl := controller.StaticController{}
    r.GET("", staticCtrl.Index)

    u := r.Group("/users")
    {
        ctrl := controller.UserController{}
        u.GET("", ctrl.Index)
        u.GET("/:id", ctrl.Show)
        u.POST("", ctrl.Create)
        u.PUT("/:id", ctrl.Update)
        u.DELETE("/:id", ctrl.Delete)
    }

    h := r.Group("/hook")
    {
    		ctrl := controller.HookController{}
            h.GET("/hello", ctrl.Index)
            h.GET("", ctrl.Verify)
            h.POST("", ctrl.Event)
    }

    return r
}
