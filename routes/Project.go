package routes

//type Routing struct {
//	routes controller.Controller
//}
//
//type RoutingInterface interface {
//	GetRoutes() *echo.Echo
//}
import (
	"github.com/labstack/echo"
)
func (Routing Routing) GetRoutesProject() *echo.Echo {
	e := echo.New()
	e.POST("/tasks/project/create/", Routing.routes.PostsCreate)
	//e.GET("/tasks/project/view/", Routing.routes.ViewProject)
	e.PUT("/tasks/project/edit/", Routing.routes.EditProject)

	return e
}