package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	dotenv "github.com/joho/godotenv"
)

func main() {
	err := dotenv.Load(".env")
	if err != nil {
		fmt.Println(".env is not loaded properly")
		os.Exit(1)
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		fmt.Println("PORT env is not loaded properly")
		os.Exit(1)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", helloHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello!\n"))
}
