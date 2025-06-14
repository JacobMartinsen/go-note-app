package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Home Page!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "About Page")
}

func initializeFileHandlers() error {
	// Disallow downloading of application
	appName, err := getAppName()
	if err != nil {
		fmt.Println("TERMINATING APP. Error getting app name:", err)
		return err
	}
	fmt.Println("App name:" + appName)
	http.HandleFunc("/"+appName, func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Forbidden.", http.StatusForbidden)
	})

	http.HandleFunc("/"+appName+".env", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Forbidden.", http.StatusForbidden)
	})

	http.Handle("/", http.FileServer(http.Dir("./")))
	return nil
}

func main() {
	godotenv.Load("app_config.env")
	fmt.Println(os.Getenv("APP_NAME"))
	// Handle routing
	http.HandleFunc("/about", aboutHandler)

	err := initializeFileHandlers()
	if err != nil {
		fmt.Println("Error initializing file handlers:", err)
		return
	}

	// Start the server
	port := getPort()
	fmt.Println("Server running on http://localhost" + port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
