package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/SERV4BIZ/coresan/config/global"
	"github.com/SERV4BIZ/gfp/jsons"

	command_datanode "github.com/SERV4BIZ/coresan/config/commands/datanode"
	command_network "github.com/SERV4BIZ/coresan/config/commands/network"
)

// WorkHandler is main job for any request
func WorkHandler(w http.ResponseWriter, r *http.Request) {
	global.MutexState.Lock()
	global.CountState++
	global.MutexState.Unlock()

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
						jsoResult = command_network.Ping(jsoCmd)
					case "datanode_info":
						jsoResult = command_datanode.Info(jsoCmd)
					case "datanode_listing":
						jsoResult = command_datanode.Listing(jsoCmd)
					}
				}
			}
		}
	}

	w.Write([]byte(jsoResult.ToString()))
}
