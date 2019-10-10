package main

import (
	"fmt"
	"github.com/e421083458/golang_common/lib"
)

func main() {
	lib.InitModule("./conf/dev/", []string{"base", "mysql", "redis"})
	defer lib.Destroy()
	//	fmt.Println()
	fmt.Println(lib.GetConfEnv())
	fmt.Println(lib.GetStringConf("base.log.log_level"))
}
