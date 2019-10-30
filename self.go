package main

import (
	"encoding/json"
	"net/http"
)

type Self struct {
	JenkinsVerison   string `json:"jenkins-version"`
	JenkinsUrl       string `json:"jenkins-url"`
	JenkinsViewCount int    `json:"jenkins-view-count"`
	JenkinsJobCount  int    `json:"jenkins-job-count"`
}

func GetSelfInfo(w http.ResponseWriter, r *http.Request) {
	var self Self
	j := Jks.JenkinsInit()
	self.JenkinsVerison = j.Version
	//获取job数量
	jobCount, _ := j.GetAllJobs()
	self.JenkinsJobCount = len(jobCount)

	viewCount, _ := j.GetAllViews()
	self.JenkinsViewCount = len(viewCount)

	//jenkins address
	self.JenkinsUrl = j.Server

	selfJson, _ := json.Marshal(self)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(selfJson)

}
