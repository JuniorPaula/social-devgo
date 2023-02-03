package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
)

func main() {
	fmt.Println("[::] web app running at 3000")

	r := router.Handler()
	log.Fatal(http.ListenAndServe(":3000", r))
}
