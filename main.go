package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	config.LoadEnv()
	utils.LoadTemplates()
	r := router.Handler()

	fmt.Printf("[::] web app running at port %d\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
