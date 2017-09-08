package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kiji/late_reason/lateReason"
)


func RegistersRoutes() {
r := gin.Default()
//lateReason/Handler
r.GET("v1/okubo/lateReason", lateReason.Handler)
r.Run(":30000")
}
