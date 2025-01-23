package main

import (
	"ADMgmtSystem/database"
	"ADMgmtSystem/library"
	"ADMgmtSystem/routers"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	router := routers.NewRouter()

	library.InitLogger()

	library.Log.Info("Service start")
	library.Log.Trace("Load env")
	err := godotenv.Load()

	if err != nil {
		library.Log.Error(err)
		log.Fatalln(err)
	}

	library.Log.Trace("Init database")
	err = database.InitDatabase()

	if err != nil {
		library.Log.Error(err)
		log.Fatalln(err)
	}

	port := os.Getenv("HTTP_PORT")
	domain := os.Getenv("HTTP_DOMAIN")
	portocol := os.Getenv("HTTP_PORTOCOL")

	library.Log.Trace("Open api service")
	library.Log.Debugf("Open url: %s://%s:%s", portocol, domain, port)

	library.OpenBrowser(fmt.Sprintf("%s://%s:%s", portocol, domain, port))

	router.Run(fmt.Sprintf("%s:%s", domain, port))
}
