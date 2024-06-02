/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"git.finsoft.id/finsoft.id/go-example/service/cmd"
)

// @title           Finsoft API
// @version         1.0
// @description     Ini adalah contoh Finsoft API.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Polma Tambunan
// @contact.url    https://www.polmatambunan.my.id/
// @contact.email  polma@finsoft.id

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8888
// @BasePath  /

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	cmd.Execute()
}
