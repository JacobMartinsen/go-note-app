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

func getPort() string {
	port := ":8080" // Default port
	if p := os.Getenv("APP_PORT_NUMBER"); p != "" {
		port = ":" + p
	}
	return port
}

func getStaticFileLocation() (string, error) {
	if dir := os.Getenv("APP_STATIC_FILES"); dir != "" {
		return "./" + dir + "/", nil
	} else {
		return "", fmt.Errorf("APP_STATIC_FILES environment variable not set")
	}
}

func main() {
	godotenv.Load("app_config.env")

	// Handle routing
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)

	// Handle static files
	staticFileLocation, err := getStaticFileLocation()
	if err != nil {
		fmt.Println("Error getting static file location:", err)
		return
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticFileLocation))))

	// Start the server
	port := getPort()
	fmt.Println("Server running on http://localhost" + port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
