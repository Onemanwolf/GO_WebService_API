package main

import (
	"fmt"
	"net/http"

	"github.com/api/webservice/controllers"
)



func main() {

	// start web server
	startWebServer(3000, 5, false)

	 controllers.RegisterControllers()
	 http.ListenAndServe(":3000", nil)
	 //panic("Error")



}



func startWebServer(port, retry int, cancel bool) (error, int) {
	fmt.Println("Starting web server")
	// do important stuff
	fmt.Println("Server started", port, retry, cancel)
	if cancel {
		fmt.Println("Cancelling")
	}
	return nil, port
}