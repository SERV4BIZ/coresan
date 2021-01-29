package storage

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/SERV4BIZ/coresan/server/global"
	"github.com/SERV4BIZ/gfp/filesystem"
	"github.com/SERV4BIZ/gfp/jsons"
)

// Read is read file from coresan
func Read(jsoCmd *jsons.JSONObject) *jsons.JSONObject {
	jsoResult := jsons.JSONObjectFactory()
	jsoResult.PutInt("status", 0)

	txtCSNID := strings.ToLower(strings.TrimSpace(jsoCmd.GetString("txt_csnid")))
	txtFullpath := global.GetFullPath(txtCSNID)
	txtInfopath := fmt.Sprint(txtFullpath, global.DS, "info.json")
	txtDatapath := fmt.Sprint(txtFullpath, global.DS, "data.dat")
	if strings.TrimSpace(txtCSNID) == "" || !filesystem.ExistFile(txtInfopath) || !filesystem.ExistFile(txtDatapath) {
		jsoResult.PutString("txt_msg", "Not exist file")
		return jsoResult
	}

	buffer, errRead := filesystem.ReadFile(txtDatapath)
	if errRead != nil {
		jsoResult.PutString("txt_msg", fmt.Sprint("Can not read file to buffer [ ", errRead, " ]"))
		return jsoResult
	}
	data := base64.StdEncoding.EncodeToString(buffer)

	jsoInfo, errInfo := jsons.JSONObjectFromFile(txtInfopath)
	if errInfo != nil {
		jsoResult.PutString("txt_msg", fmt.Sprint("Can not parse json info from file [ ", errInfo, " ]"))
		return jsoResult
	}
	jsoInfo.PutString("txt_data", data)

	jsoResult.PutObject("jso_data", jsoInfo)
	jsoResult.PutInt("status", 1)
	return jsoResult
}
