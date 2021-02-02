package drivers

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/SERV4BIZ/coresan/api/networks"
	"github.com/SERV4BIZ/gfp/jsons"
)

// Read is read file from coresan
func Read(jsoHost *jsons.JSONObject, txtCSNID string) (*jsons.JSONObject, []byte) {
	jsoCmd := jsons.JSONObjectFactory()
	jsoCmd.PutString("txt_command", "storage_read")
	jsoCmd.PutString("txt_csnid", strings.ToLower(strings.TrimSpace(txtCSNID)))

	jsoAuthen := jsons.JSONObjectFactory()
	jsoAuthen.PutString("txt_username", strings.TrimSpace(strings.ToLower(jsoHost.GetString("txt_username"))))
	jsoAuthen.PutString("txt_password", strings.TrimSpace(jsoHost.GetString("txt_password")))
	jsoCmd.PutObject("jso_authen", jsoAuthen)

	jsoReq := networks.Request(jsoHost, jsoCmd)
	if jsoReq.GetInt("status") > 0 {
		jsoData := jsoReq.GetObject("jso_data")
		buffer, errBase64 := base64.StdEncoding.DecodeString(jsoData.GetString("txt_data"))
		if errBase64 != nil {
			jsoResult := jsons.JSONObjectFactory()
			jsoResult.PutInt("status", 0)
			jsoResult.PutString("txt_msg", fmt.Sprint("Can not decode base64 of data [ ", errBase64, " ]"))
			return jsoResult, nil
		}

		jsoData.Remove("txt_data")
		return jsoReq, buffer
	}

	return jsoReq, nil
}
