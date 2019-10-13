package controller

import (
	"encoding/json"
	"github.com/bndr/gojenkins"
	"github.com/gin-gonic/gin"
	"github.com/msupers/hobbits/middleware"
)

type Jenkins struct {
}

func JenkinsRegister(router *gin.RouterGroup) {
	jenkins := Jenkins{}
	router.GET("/index", jenkins.Index)
	router.GET("/jobs", jenkins.JobList)
	//router.GET("/bind", demo.Bind)
	//router.GET("/dao", demo.Dao)
	//router.GET("/redis", demo.Redis)
}

func (jenkins *Jenkins) Index(c *gin.Context) {
	middleware.ResponseSuccess(c, "")
	return
}
func (jenkins *Jenkins) JobList(c *gin.Context) {
	//jenkins init
	//jenkins := gojenkins.CreateJenkins(nil, "http://localhost:8080/", "admin", "admin")
	j := gojenkins.CreateJenkins("http://10.2.40.44:8080/", "admin", "123.com").Init()

	type JobList struct {
		JobName string
		Color   string
	}

	jl := j.GetAllJobs(true)
	var Jbl []JobList
	for i, _ := range jl {
		var j JobList

		j.JobName = jl[i].Raw.Name
		j.Color = jl[i].Raw.Color
		//Color:jl[i].Raw.Color}}
		Jbl = append(Jbl, j)
		//aa :=
		//jbl = append(jbl,aa)
		//fmt.Printf("job name is %v job desc is %v\n",jobList[i].Raw.Name,jobList[i].Raw.Description)
		//	jobList[i].Raw.Description
	}
	jm, _ := json.Marshal(Jbl)
	middleware.ResponseSuccess(c, string(jm))
	//fmt.Println(jobList[0].Jenkins)
	//jl := j.GetJob("new-test")
	/*
		jl := j.GetJob("new-test")
		var res struct {
			Jenkins string
			Base    string
		}
		res.Jenkins = jl.Jenkins.Server
		res.Base = jl.Base
		resp, err := json.Marshal(res)
		if err != nil {
			middleware.ResponseError(c, 400, err)
		}

		middleware.ResponseSuccess(c, string(resp))
	*/
}
