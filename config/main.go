package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/SERV4BIZ/coresan/config/global"
	"github.com/SERV4BIZ/coresan/config/locals"
	"github.com/SERV4BIZ/coresan/config/utility"
)

func main() {
	jsoConfig, errConfig := locals.LoadConfig()
	if errConfig != nil {
		panic(errConfig)
	}
	global.JSOConfig = jsoConfig

	fmt.Println("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * *")
	fmt.Println(fmt.Sprint(global.AppName, " Version ", global.AppVersion))
	fmt.Println(fmt.Sprint("Copyright © 2020 ", global.CompanyName, " All Rights Reserved."))
	fmt.Println("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * *")
	fmt.Println(fmt.Sprint("Directory : ", utility.GetAppDir()))
	fmt.Println("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * *")
	fmt.Println("Loading configuration file.")
	fmt.Println("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * *")
	fmt.Println("")
	fmt.Println(global.JSOConfig.ToString())
	fmt.Println("")
	fmt.Println("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * *")

	go func() {
		for {
			var m runtime.MemStats
			runtime.ReadMemStats(&m)

			global.MutexState.Lock()
			global.MemoryState = int(utility.NumberByteToMb(m.Sys))
			global.LoadState = global.CountState
			global.CountState = 0
			global.MutexState.Unlock()

			<-time.After(time.Second)
		}
	}()

	intTime := global.JSOConfig.GetInt("int_timeout")
	if intTime <= 0 {
		intTime = 15
	}

	global.Username = global.JSOConfig.GetString("txt_username")
	global.Password = global.JSOConfig.GetString("txt_password")

	router := http.NewServeMux()
	router.HandleFunc("/", WorkHandler)

	appsrv := &http.Server{
		Addr:         fmt.Sprint(":", global.JSOConfig.GetInt("int_port")),
		Handler:      router,
		ReadTimeout:  time.Duration(intTime) * time.Second,
		WriteTimeout: time.Duration(intTime) * time.Second,
	}
	appsrv.ListenAndServe()
}
