package main

import (
	"encoding/json"
	"fmt"
	"log"
	"main/api"
	"main/config"
	"main/job"
	"main/logger"
	"main/pkg/data"
	"net/http"
)

//Basic middware to log requests.
func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headerString, _ := json.Marshal(r.Header)
		logger.Log.Printf("[Req Recv] path:%s | headers:%s", r.URL, headerString)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	/*
		Init operations
	 */
	data.LoadFromFile()
	job.TimedJobStart()

	/**
		Http server start operation
	 */
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		res.Write([]byte(""))
	})

	mux.HandleFunc("/set", api.Set)
	mux.HandleFunc("/get", api.Get)
	mux.HandleFunc("/flush", api.Flush)

	logger.Log.Printf("Server listening port:%s", config.SERVICE_PORT)
	server := http.Server{Addr: fmt.Sprintf(":%s", config.SERVICE_PORT), Handler: logRequest(mux)}
	log.Fatal(server.ListenAndServe())
}

