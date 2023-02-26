package main

import (
	"MyProject/config"
	"MyProject/router"
	"fmt"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	log.SetPrefix("------------------ ")
	config.SetUp()

	r := router.InitRouter()

	serverPort := fmt.Sprintf(":%s", config.SERVER_PORT)
	r.Run(serverPort)
}
