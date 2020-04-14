package restful

import (
	"github.com/gin-gonic/gin"
	"github.com/ontio/sagapi/restful/api"
	"github.com/ontio/sagapi/middleware/jwt"
)

func NewRouter() *gin.Engine {
	root := gin.New()
	root.Use(jwt.JWT())
	api.RoutesApi(root)
	return root
}
