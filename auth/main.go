package main

import (
	"github.com/architecture-template/echo-ddd/auth/presentation/router"
)

// @title Chat Connect
// @version 1.0
// @description This is a sample swagger server.
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8002
// @BasePath /
func main() {
	router.Init()
}
