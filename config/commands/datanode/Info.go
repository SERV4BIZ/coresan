package datanode

import (
	"fmt"
	"strings"

	"github.com/SERV4BIZ/coresan/config/locals"
	"github.com/SERV4BIZ/gfp/jsons"
)

// Info is get info of datanode
func Info(jsoCmd *jsons.JSONObject) *jsons.JSONObject {
	jsoResult := jsons.JSONObjectFactory()
	jsoResult.PutInt("status", 0)

	nodeName := strings.TrimSpace(strings.ToLower(jsoCmd.GetString("txt_name")))
	nodeInfo, errNodeInfo := locals.LoadDataNodeInfo(nodeName)
	if errNodeInfo != nil {
		jsoResult.PutString("txt_msg", fmt.Sprint("Can not load data node info [ ", errNodeInfo, " ]"))
		return jsoResult
	}

	jsoResult.PutObject("jso_data", nodeInfo)
	jsoResult.PutInt("status", 1)
	return jsoResult
}
