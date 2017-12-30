package echo_restful_resource

import "github.com/labstack/echo"

func New(e echo.Echo, prefix string, resource Resource) (group echo.Group) {
	g := e.Group(prefix)
	g.GET("", resource.Index)
	g.GET("/create", resource.Create)
	g.POST("", resource.Store)
	g.GET("/:id", resource.Show)
	g.GET("/:id/edit", resource.Edit)
	g.PUT("/:id", resource.Update)
	g.DELETE("/:id", resource.Destroy)
}


