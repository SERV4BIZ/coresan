package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/SERV4BIZ/coresan/server/global"
	"github.com/SERV4BIZ/coresan/server/locals"
	"github.com/SERV4BIZ/coresan/server/utility"
	"github.com/SERV4BIZ/gfp/files"
)

func main() {
	var errConfig error
	global.JSOConfig, errConfig = locals.LoadConfig()
	if errConfig != nil {
		panic(errConfig)
	}
	global.NFSPath = global.JSOConfig.GetString("txt_path_nfs")
	global.Username = global.JSOConfig.GetString("txt_username")
	global.Password = global.JSOConfig.GetString("txt_password")

	global.MaxRead = global.JSOConfig.GetInt("int_maxread")
	if global.MaxRead <= 0 {
		// Default max reader is 1024MB or 1GB
		global.MaxRead = 1024 * 1024 * 1024
	}

	fmt.Println("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * *")
	fmt.Println(fmt.Sprint(global.AppName, " Version ", global.AppVersion))
	fmt.Println("Copyright © 2019 Serv4Biz Co.,Ltd. All Rights Reserved.")
	fmt.Println("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * *")
	fmt.Println(fmt.Sprint("Directory : ", utility.GetAppDir()))
	fmt.Println("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * *")
	fmt.Println("Loading configuration file.")
	fmt.Println("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * *")
	fmt.Println("")
	fmt.Println(global.JSOConfig.ToString())
	fmt.Println("")
	fmt.Println("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * *")
	files.MakeDir(global.NFSPath)

	// Load and Memory Monitor
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

	// Force GC to clear up
	go func() {
		for {
			<-time.After(time.Hour)
			runtime.GC()
		}
	}()

	intTime := global.JSOConfig.GetInt("int_timeout")
	if intTime <= 0 {
		intTime = 15
	}

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
