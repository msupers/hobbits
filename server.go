package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	//some job handle func
	mux.HandleFunc("/jenkins/job", JenkinsJob)
	mux.HandleFunc("/self/info", GetSelfInfo)
	//mux.HandleFunc("/job/list", ListJenkinsJob)
	mux.HandleFunc("/job/config", GetJobConfig)

	//some view handle func
	//	mux.HandleFunc("/view/list",ListJenkinsView)
	server := &http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: mux,
	}
	err := server.ListenAndServe()
	log.Fatal(err)
}
