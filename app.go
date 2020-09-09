package main

import "plm/routes"

// Starting server
func main() {
	echo := routes.Routing.GetRoutes(routes.Routing{})
	//echo := routes.Routing.GetRoutesProject(routes.Routing{})
	_ = echo.Start(":9001")
}
