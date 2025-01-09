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

	err := godotenv.Load()

	if err != nil {
		log.Fatalln(err)
	}

	err = database.InitDatabase()

	if err != nil {
		log.Fatalln(err)
	}

	port := os.Getenv("HTTP_PORT")
	domain := os.Getenv("HTTP_DOMAIN")
	portocol := os.Getenv("HTTP_PORTOCOL")

	library.OpenBrowser(fmt.Sprintf("%s://%s:%s", portocol, domain, port))

	router.Run(fmt.Sprintf("%s:%s", domain, port))
}
