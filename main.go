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
