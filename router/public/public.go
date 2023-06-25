package public

import (
	"github.com/gin-gonic/gin"
)

type PublicRouterGroup struct {
}

func (p *PublicRouterGroup) AddPublicRouterGroup(g *gin.RouterGroup) {
	g.GET("health", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "ok",
		})
	})

	// userApi := v1.AppApiGroupIns.SysApiGroup.UserApi
	// g.POST("register", userApi.Register)
}
