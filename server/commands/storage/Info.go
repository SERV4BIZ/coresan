package storage

import (
	"fmt"
	"strings"

	"github.com/SERV4BIZ/coresan/server/global"
	"github.com/SERV4BIZ/gfp/files"
	"github.com/SERV4BIZ/gfp/jsons"
)

// Info is get info file
func Info(jsoCmd *jsons.JSONObject) *jsons.JSONObject {
	jsoResult := jsons.JSONObjectFactory()
	jsoResult.PutInt("status", 0)

	txtCSNID := strings.ToLower(strings.TrimSpace(jsoCmd.GetString("txt_csnid")))
	txtFullpath := global.GetFullPath(txtCSNID)
	txtInfopath := fmt.Sprint(txtFullpath, global.DS, "info.json")
	txtDatapath := fmt.Sprint(txtFullpath, global.DS, "data.dat")
	if strings.TrimSpace(txtCSNID) == "" || !files.ExistFile(txtInfopath) || !files.ExistFile(txtDatapath) {
		jsoResult.PutString("txt_msg", "Not exist file")
		return jsoResult
	}

	jsoInfo, errInfo := jsons.JSONObjectFromFile(txtInfopath)
	if errInfo != nil {
		jsoResult.PutString("txt_msg", fmt.Sprint("Can not parse json info from file [ ", errInfo, " ]"))
		return jsoResult
	}

	jsoResult.PutObject("jso_data", jsoInfo)
	jsoResult.PutInt("status", 1)
	return jsoResult
}
