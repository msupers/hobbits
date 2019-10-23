package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	//mux.HandleFunc("/",index)
	//mux.HandleFunc("/",index)
	//mux.HandleFunc("/job/create",createJenkinsJob)
	mux.HandleFunc("/job/run", RunJenkinsJob)
	mux.HandleFunc("/job/list", ListJenkinsJob)
	server := &http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: mux,
	}
	server.ListenAndServe()
}
