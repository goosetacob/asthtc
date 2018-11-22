package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/goosetacob/asthtc/backend/resource"
	"github.com/gorilla/mux"
)

var port string

func init() {
	// configure logrus timestamp
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	logrus.SetFormatter(customFormatter)

	// configure port address
	portAddress := os.Getenv("PORT")
	switch len(port) {
	case 0:
		logrus.Println("Environment variable PORT is undefined. Using port :80 by default")
		port = ":80"
	case 1:
		logrus.Printf("Environment variable PORT=\"%s\"", port)
		port = ":" + portAddress
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		logrus.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/tool/voweless", func(w http.ResponseWriter, r *http.Request) {
		type vowelessJob struct {
			Phrase string `json:"phrase"`
		}

		var job vowelessJob
		var err error
		if err = json.NewDecoder(r.Body).Decode(&job); err != nil {
			logrus.Error(err)
			fmt.Fprintf(w, "%v", err)
		}

		var response string
		if response, err = tool.Voweless(job.Phrase); err != nil {
			logrus.Error(err)
			fmt.Fprintf(w, "%v", err)
		}

		fmt.Fprintf(w, "Phrase: %v\nResponse: %v\n", job.Phrase, response)
	}).Methods("POST")

	r.HandleFunc("/tool/aesthetic", func(w http.ResponseWriter, r *http.Request) {
		type aestheticJob struct {
			Phrase string `json:"phrase"`
		}

		var job aestheticJob
		var err error
		if err = json.NewDecoder(r.Body).Decode(&job); err != nil {
			logrus.Error(err)
			fmt.Fprintf(w, "%v", err)
		}

		var response string
		if response, err = tool.Aesthetic(job.Phrase); err != nil {
			logrus.Error(err)
			fmt.Fprintf(w, "%v", err)
		}

		fmt.Fprintf(w, "Phrase: %v\nResponse: %v\n", job.Phrase, response)
	}).Methods("POST")

	r.HandleFunc("/tool/debruijn", func(w http.ResponseWriter, r *http.Request) {
		type debruijnJob struct {
			Alphabet        string `json:"alphabet"`
			SubSequenceSize int    `json:"subSequenceSize"`
		}

		var job debruijnJob
		var err error
		if err = json.NewDecoder(r.Body).Decode(&job); err != nil {
			logrus.Error(err)
			fmt.Fprintf(w, "%v", err)
		}

		var response string
		if response, err = tool.DeBruijn(job.Alphabet, job.SubSequenceSize); err != nil {
			logrus.Error(err)
			fmt.Fprintf(w, "%v", err)
		}

		fmt.Fprintf(w, "Alphabet: %v\nSubSequenceSize: %v\nResponse: %v\n", job.Alphabet, job.SubSequenceSize, response)
	}).Methods("POST")

	r.Use(loggingMiddleware)
	http.ListenAndServe(port, r)
}
