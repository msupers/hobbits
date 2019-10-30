package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bndr/gojenkins"
	"html/template"
	"net/http"
	"os"
)

type Jenkins struct{}

var Jks Jenkins

func (jenkins *Jenkins) JenkinsInit() *gojenkins.Jenkins {
	j, _ := gojenkins.CreateJenkins(nil, "http://10.2.40.44:8080/", "admin", "123.com").Init()
	return j
	//fmt.Printf("%T", jks)
}

type JobInfo struct {
	JobName    string `json:"jobname"`
	Color      string `json:"color"`
	BuiltCount int64  `json:"builtcount"`
}

var jobInfo JobInfo
var jobList []JobInfo

func JenkinsJob(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = GetJenkinsJob(w, r)
	case "POST":
		err = PostJenkinsJob(w, r)
	case "PUT":
		err = PutJenkinsJob(w, r)
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
	jobName := r.FormValue("name")
	if jobName == "" {
		allJob, _ := j.GetAllJobs()
		for i := range allJob {
			jobInfo.JobName = allJob[i].Raw.Name
			jobInfo.Color = allJob[i].Raw.Color
			jobInfo.BuiltCount = allJob[i].Raw.LastBuild.Number
			jobList = append(jobList, jobInfo)
		}
		jobListJson, err := json.Marshal(jobList)
		if err != nil {
			return err
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(jobListJson)

	} else {
		oneJob, _ := j.GetJob(jobName)
		jobInfo.JobName = oneJob.Raw.Name
		jobInfo.Color = oneJob.Raw.Color
		jobInfo.BuiltCount = oneJob.Raw.LastBuild.Number
		oneJobJson, err := json.Marshal(jobInfo)
		if err != nil {
			return err
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(oneJobJson)
	}

	return
}

func PostJenkinsJob(w http.ResponseWriter, r *http.Request) (err error) {
	type TplJob struct {
		GitProject string
		GitGroup   string
	}
	j := Jks.JenkinsInit()
	var tpl TplJob
	gitProject := r.FormValue("project")
	gitGroup := r.FormValue("group")
	tpl.GitProject = gitProject
	tpl.GitGroup = gitGroup
	jobName := gitGroup + "-" + gitProject + "-master"

	//生成jenkins job config 前半部分
	pipelineHead, err := os.Open("templates/jenkins_tpl_head.xml")
	if err != nil {
		fmt.Println(err)
	}
	buf := make([]byte, 20000)
	n, _ := pipelineHead.Read(buf)
	strHead := string(buf[:n])
	defer pipelineHead.Close()

	//生成jenkins Job config 后半部分
	pipelineTail, err := os.Open("templates/jenkins_tpl_pipeline.groovy")
	if err != nil {
		fmt.Println(err)
	}
	buf2 := make([]byte, 20000)
	tail, _ := pipelineTail.Read(buf2)
	strTail := string(buf2[:tail])
	defer pipelineTail.Close()
	//fmt.Printf("%T", str)
	textTpl, err := template.New("test").Parse(strHead + "\n" + strTail)
	if err != nil {
		return err
	}
	var b1 bytes.Buffer
	_ = textTpl.Execute(&b1, tpl)
	result, err := j.CreateJob(b1.String(), jobName)
	if err != nil {
		return err
	}
	_, _ = fmt.Fprintln(w, result)
	return nil
}

func PutJenkinsJob(w http.ResponseWriter, r *http.Request) (err error) {
	j := Jks.JenkinsInit()
	jobName := r.FormValue("name")
	if jobName == "" {
		_, _ = fmt.Fprintln(w, "jobname can not null")
		return
	}
	startNo, err := j.BuildJob(jobName)
	if err != nil {
		return err
	}
	jobStartResult := struct {
		JobName string `json:"job-name"`
		IsStart bool   `json:"is-start"`
		StartNo int64  `json:"start-no"`
	}{JobName: jobName, IsStart: true, StartNo: startNo}

	jobStartJson, err := json.Marshal(jobStartResult)
	_, _ = w.Write(jobStartJson)
	//_, _ = fmt.Fprintf(w, jobName+" is starting running, start No is %d", startNo)

	//w.Write()
	return err
}
func DeleteJenkinsJob(w http.ResponseWriter, r *http.Request) (err error) {
	//fmt.Fprintln(w, "delete jenkins job ", http.StatusOK)
	j := Jks.JenkinsInit()
	jobname := r.FormValue("jobname")
	delResult, err := j.DeleteJob(jobname)
	_, _ = fmt.Fprintln(w, delResult)
	return err
}

//GET /job/config?jobname=$JobName
// get jenkins job config
func GetJobConfig(w http.ResponseWriter, r *http.Request) {
	j := Jks.JenkinsInit()
	jobname := r.FormValue("jobname")
	job, _ := j.GetJob(jobname)
	config, _ := job.GetConfig()
	_, _ = fmt.Fprintln(w, config)
}
