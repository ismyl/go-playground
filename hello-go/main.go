package main

import (
	"fmt"
	"github.com/ismyl/webservice/controllers"
	"net/http"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(`:3000`, nil)
}

func startWebServer(port, numberOfRetries int) (int, error) {
	fmt.Println("Starting server...")
	// do Some server running code
	fmt.Println("Server started on port", port)
	fmt.Println("Number of retries", numberOfRetries)

	// return nil
	return port, nil
}
