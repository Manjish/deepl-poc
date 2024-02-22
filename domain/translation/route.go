package translation

import (
  "github.com/gin-gonic/gin"
  "deepl-poc/pkg/infrastructure"
)

type Route struct {
    router *infrastructure.Router
    controller *Controller
    groupRouter *gin.RouterGroup
}

func NewRoute(router *infrastructure.Router, controller *Controller) *Route {
	route := Route{router: router, controller: controller}
  route.groupRouter = router.Group("api/translation")
	return &route
}

func RegisterRoute(r *Route) {
	r.groupRouter.GET("", r.controller.HandleRoot)
	r.groupRouter.POST("/translate", r.controller.GetTransalation)
}
