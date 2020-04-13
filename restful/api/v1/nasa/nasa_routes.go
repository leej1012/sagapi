package nasa

import (
	"github.com/gin-gonic/gin"
)

func RoutesNasa(parent *gin.RouterGroup) {
	nasaRouteGroup := parent.Group("/nasa")
	nasaRouteGroup.GET("/apod/:apikey", Apod)
	nasaRouteGroup.GET("/feed/:startdate/:enddate/:apikey", Feed)
}
