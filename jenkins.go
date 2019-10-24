package main

import (
	"encoding/json"
	"fmt"
	"github.com/bndr/gojenkins"
	"net/http"
)

type Jenkins struct{}

var Jks Jenkins

func (jenkins *Jenkins) JenkinsInit() *gojenkins.Jenkins {
	j, _ := gojenkins.CreateJenkins(nil, "http://10.2.40.43:8080/", "admin", "123.com").Init()
	return j
	//fmt.Printf("%T", jks)
}

type JobInfo struct {
	JobName    string `json:"jobname"`
	Color      string `json:"color"`
	BuiltCount int64  `json:"builtcount"`
}

var jInfo JobInfo
var jList []JobInfo

func JenkinsJob(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = GetJenkinsJob(w, r)
	case "POST":
		err = CreateJenkinsJob(w, r)
	case "PUT":
		err = RunJenkinsJob(w, r)
	case "DELETE":
		err = DeleteJenkinsJob(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//GET /jenkins/job
func GetJenkinsJob(w http.ResponseWriter, r *http.Request) (err error) {
	j := Jks.JenkinsInit()
	jobname := r.FormValue("jobname")
	if jobname == "" {
		allJob, _ := j.GetAllJobs()
		for i := range allJob {
			jInfo.JobName = allJob[i].Raw.Name
			jInfo.Color = allJob[i].Raw.Color
			jInfo.BuiltCount = allJob[i].Raw.LastBuild.Number
			jList = append(jList, jInfo)
		}
		json, _ := json.Marshal(jList)
		fmt.Fprintln(w, string(json))
		return nil
		//fmt.Fprintln(w,)
	} else {
		oneJob, _ := j.GetJob(jobname)
		jInfo.JobName = oneJob.Raw.Name
		jInfo.Color = oneJob.Raw.Color
		jInfo.BuiltCount = oneJob.Raw.LastBuild.Number
		json, err := json.Marshal(jInfo)
		fmt.Fprintln(w, string(json))
		return err
	}

	return nil
}

func CreateJenkinsJob(w http.ResponseWriter, r *http.Request) (err error) {
	fmt.Fprintln(w, "create jenkins job ", http.StatusOK)
	return nil
}

func RunJenkinsJob(w http.ResponseWriter, r *http.Request) (err error) {
	fmt.Fprintln(w, "run jenkins job ", http.StatusOK)
	return nil
}
func DeleteJenkinsJob(w http.ResponseWriter, r *http.Request) (err error) {
	fmt.Fprintln(w, "delete jenkins job ", http.StatusOK)
	return nil
}

/*
// PUT /job/run?jobname=$JobName
// run a jenkins job by job name
func RunJenkinsJob(w http.ResponseWriter, r *http.Request) {
	j := Jks.JenkinsInit()
	jobname := r.FormValue("jobname")
	result, err := j.BuildJob(jobname)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Fprintln(w, result)
}

//GET /job/list
// get all job list
type JobList struct {
	JobName    string `json:"jobname"`
	Color      string `json:"color"`
	BuiltCount int64  `json:"builtcount"`
}

func ListJenkinsJob(w http.ResponseWriter, r *http.Request) {
	var jobInfo []JobList
	var jL JobList
	j := Jks.JenkinsInit()
	jobList, _ := j.GetAllJobs()
	for i := range jobList {
		jL.Color = jobList[i].Raw.Color
		jL.JobName = jobList[i].Raw.Name
		jL.BuiltCount = jobList[i].Raw.LastBuild.Number
		jobInfo = append(jobInfo, jL)
	}
	json, err := json.Marshal(jobInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Fprintln(w, string(json))
}

//GET /job/config?jobname=$JobName
// get jenkins job config
func GetJobConfig(w http.ResponseWriter, r *http.Request) {
	j := Jks.JenkinsInit()
	jobname := r.FormValue("jobname")
	job, _ := j.GetJob(jobname)
	config, _ := job.GetConfig()
	fmt.Fprintln(w, config)
}

type ViewList struct {
	ViewName string `json:"view-name"`
}

func ListJenkinsView(w http.ResponseWriter, r *http.Request) {
	var vList ViewList
	var vListRespon []ViewList
	j := Jks.JenkinsInit()
	viewList, _ := j.GetAllViews()
	for i := range viewList {
		vList.ViewName = viewList[i].Raw.Name
		vListRespon = append(vListRespon, vList)
	}
	json, _ := json.Marshal(vListRespon)

	fmt.Fprintln(w, string(json))
}
*/
