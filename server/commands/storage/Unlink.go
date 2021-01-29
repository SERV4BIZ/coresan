package storage

import (
	"fmt"
	"strings"

	"github.com/SERV4BIZ/coresan/server/global"
	"github.com/SERV4BIZ/gfp/files"
	"github.com/SERV4BIZ/gfp/jsons"
)

// Unlink is delete file from coresan
func Unlink(jsoCmd *jsons.JSONObject) *jsons.JSONObject {
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

	errDel := files.DeleteFile(txtFullpath)
	if errDel != nil {
		jsoResult.PutString("txt_msg", fmt.Sprint("Can not delete file [ ", errDel, " ]"))
		return jsoResult
	}

	jsoResult.PutInt("status", 1)
	return jsoResult
}
