package routers

import (
    "github.com/gin-gonic/gin"
    "monitor/routers/api/v1"
    "monitor/setting"
)

func InitRouter() *gin.Engine{
    r := gin.New()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())
    gin.SetMode(setting.RunMode)
    apiv1 := r.Group("/api/v1")
    {
        apiv1.GET("/system/cpu",v1.GetCpu)
        apiv1.GET("/system/mem",v1.GetMem)
        apiv1.GET("/system/io",v1.GetIo)
        apiv1.GET("/system/net",v1.GetNet)

    }

    return r
}
