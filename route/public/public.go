package public

import (
	"github.com/gin-gonic/gin"
)

type AppPublicRouteGroup struct {
}

func (p *AppPublicRouteGroup) AddPublicGroup(g *gin.RouterGroup) {
	g.GET("health", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "ok",
		})
	})

	// userApi := v1.AppApiGroupIns.SysApiGroup.UserApi
	// g.POST("register", userApi.Register)
}
