package main

import (
	"github.com/nmpotential/webservice/4-creating-functions-and-methods/8-starting-the-webservice/controllers"
	"net/http"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(" :3000", nil)
}