package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"

	dotenv "github.com/joho/godotenv"
)

func main() {
	err := dotenv.Load(".env")
	if err != nil {
		fmt.Println(".env is not loaded properly")
		os.Exit(1)
	}

	// Configure Logging
	logFileLocation := os.Getenv("LOG_FILE_LOCATION")
	if logFileLocation != "" {
		log.SetOutput(&lumberjack.Logger{
			Filename:   logFileLocation,
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     28,   //days
			Compress:   true, // disabled by default
		})
	}

	log.Println("env loaded")

	port, ok := os.LookupEnv("PORT")
	if !ok {
		fmt.Println("PORT env is not loaded properly")
		os.Exit(1)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", helloHandler)

	log.Println("server up")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("handling %s %s %s\n", r.Method, r.RequestURI, r.RemoteAddr)
	w.Write([]byte("hello!\n"))
}
