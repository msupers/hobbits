package main

import (
	"fmt"
	"github.com/msupers/hobbits/public"
	"github.com/msupers/hobbits/router"
	"os"
	"os/signal"
	"syscall"

	"github.com/e421083458/golang_common/lib"
)

func main() {
	lib.InitModule("./conf/dev/", []string{"base", "mysql", "redis"})
	defer lib.Destroy()
	//	fmt.Println()
	//fmt.Println(lib.GetConfEnv())
	fmt.Println(lib.GetStringConf("base.log.log_level"))
	fmt.Println(lib.GetStringConf("base.base.time_location"))
	//fmt.Println(lib.GetStringConf("base.jenkins.url"))
	fmt.Println(lib.GetStringConf("base.jenkins.username"))

	public.InitMysql()
	public.InitValidate()
	router.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
}
