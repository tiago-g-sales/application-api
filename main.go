package main

import (
	"github.com/example/application-api/src/application/controllers"
	"github.com/colibri-project-io/colibri-sdk-go"
	"github.com/colibri-project-io/colibri-sdk-go/pkg/web/restserver"
)

func init() {
	colibri.InitializeApp()
	//storage.Initialize() // uncomment if you use storage
	//cacheDB.Initialize() // uncomment if you use cache
	//sqlDB.Initialize() // uncomment if you use sql database
	//messaging.Initialize() // uncomment if you use messaging
	// application-api
}

func main() {
	restserver.AddRoutes(controllers.NewDemoController().Routes())
	restserver.ListenAndServe()
}
