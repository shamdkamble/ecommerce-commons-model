package server

import (
	"fmt"
	"net/http"

	"../commons/utils"
)

func StartServer() {
	fmt.Println("Ecommerce server is running...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Verify request
		if utils.IsRequestValid(r) {
			// Redirect to db server
			// Handle response from db
			// Write response to client
		} else {
			fmt.Println("Invalid request")
			// Write error response to client
		}
	})
	http.ListenAndServe(":3100", nil)
}
