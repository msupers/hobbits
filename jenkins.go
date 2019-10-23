package main

import (
	"fmt"
	"github.com/bndr/gojenkins"
	"net/http"
)

type Jenkins struct{}

var Jks Jenkins

func (jenkins *Jenkins) JenkinsInit() *Jenkins {
	j := CreateJenkins("http://10.2.40.43:8080/", "admin", "123.com").Init()
	return j
	//fmt.Printf("%T", jks)
}

func RunJenkinsJob(w http.ResponseWriter, r *http.Request) {
	j := Jks.JenkinsInit()
	jobname := r.FormValue("jobname")
	result := j.BuildJob(jobname)
	////fmt.Println(jobname)
	job := j.GetJob(jobname)
	////fmt.Println(job.Raw.Name)
	//aa := job.GetName()
	fmt.Println(job.Raw.Builds)
	//j.Poll()
	//j.BuildJob(job.Raw.Name)
	//
	//fmt.Println(aa)
	//result,err := j.BuildJob(jobname)
	//fmt.Println(result)
	if result {
		fmt.Fprintln(w, jobname+" is running !")
	} else {
		fmt.Fprintln(w, jobname+" not run running")
	}

}

func ListJenkinsJob(w http.ResponseWriter, r *http.Request) {
	j := Jks.JenkinsInit()
	aaa := j.GetAllJobs(true)
	//fmt.Fprintln(w, aaa)
	for i := range aaa {
		fmt.Println(aaa[i].Raw.Name)
	}
}

//func (Jenkins *Jenkins)createJenkinsJob(w http.ResponseWriter, r *http.Request) {
//	j := Jenkins.JenkinsInit()
//
//
//}
//jenkinsInit()
