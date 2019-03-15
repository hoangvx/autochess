package server

import (
    "github.com/gin-gonic/gin"

    "autochess/controller"
)

// Init is initialize server
func Init() {
    r := router()
    r.Run()
}

func router() *gin.Engine {
    r := gin.Default()
    r.LoadHTMLGlob("client/*.html")
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
