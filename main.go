package main

import (
	"net/http"

	"github.com/pluralsight/webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}

// func startWebServer(port, numberOfRetries int) (int, error) {
// 	fmt.Println("Starting server...")

// 	fmt.Println("Server started on port", port)
// 	fmt.Println("Number of retries", numberOfRetries)
// 	return port, nil
// }
