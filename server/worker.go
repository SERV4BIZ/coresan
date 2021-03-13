package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/SERV4BIZ/coresan/server/commands/network"
	"github.com/SERV4BIZ/coresan/server/commands/storage"
	"github.com/SERV4BIZ/coresan/server/global"
	"github.com/SERV4BIZ/gfp/jsons"
)

// WorkHandler is main handler all command
func WorkHandler(w http.ResponseWriter, r *http.Request) {
	global.MutexState.Lock()
	global.CountState++
	global.MutexState.Unlock()

	// Max memory 1024MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(global.MaxRead))
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	jsoResult := jsons.JSONObjectFactory()
	jsoResult.PutInt("status", 0)

	buffer, errBody := ioutil.ReadAll(r.Body)
	if errBody != nil {
		jsoResult.PutString("txt_msg", fmt.Sprint("Can not read body from http request [ ", errBody, " ]"))
	} else {
		jsoCmd, errCmd := jsons.JSONObjectFromString(string(buffer))
		if errCmd != nil {
			jsoResult.PutString("txt_msg", fmt.Sprint("Can not load command from json string buffer [ ", errCmd, " ]"))
		} else {
			jsoAuthen := jsoCmd.GetObject("jso_authen")
			authenUser := strings.TrimSpace(strings.ToLower(jsoAuthen.GetString("txt_username")))
			authenPass := strings.TrimSpace(jsoAuthen.GetString("txt_password"))

			if len(authenUser) > 0 {
				if global.Username == authenUser && global.Password == authenPass {
					switch jsoCmd.GetString("txt_command") {
					case "network_ping":
						jsoResult = network.Ping(jsoCmd)
					case "storage_exist":
						jsoResult = storage.Exist(jsoCmd)
					case "storage_info":
						jsoResult = storage.Info(jsoCmd)
					case "storage_read":
						jsoResult = storage.Read(jsoCmd)
					case "storage_write":
						jsoResult = storage.Write(jsoCmd)
					case "storage_rewrite":
						jsoResult = storage.Rewrite(jsoCmd)
					case "storage_unlink":
						jsoResult = storage.Unlink(jsoCmd)
					}
				}
			}
		}
	}
	w.Write([]byte(jsoResult.ToString()))
}
