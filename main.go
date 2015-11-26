package main

import (
	"gojira/routers"
)

func main() {

	router := routers.InitRoutes()

	router.Run(":8081")

}
