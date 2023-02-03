package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("[::] web app running at 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
