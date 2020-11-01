package main

import (
	"github.com/wrighbr/resume-api/api"
)

// @title Resume API
// @version 1.0
// @description REST API for my resume
// @securityDefinitions.basic BasicAuth
// @in header
// @name Authorization
// @contact.name Brett Wright
// @contact.email wright.brett@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /

func main() {
	api.HandleRequests()

}
