package drivers

import (
	"encoding/base64"
	"strings"

	"github.com/SERV4BIZ/coresan/api/networks"
	"github.com/SERV4BIZ/gfp/jsons"
)

// Write is write file to coresan
func Write(jsoHost *jsons.JSONObject, txtFilename string, dblExpire float64, buffer []byte) *jsons.JSONObject {
	jsoCmd := jsons.JSONObjectFactory()
	jsoCmd.PutString("txt_command", "storage_write")
	jsoCmd.PutString("txt_filename", txtFilename)
	jsoCmd.PutDouble("dbl_expire", dblExpire)
	jsoCmd.PutString("txt_data", base64.StdEncoding.EncodeToString(buffer))

	jsoAuthen := jsons.JSONObjectFactory()
	jsoAuthen.PutString("txt_username", strings.TrimSpace(strings.ToLower(jsoHost.GetString("txt_username"))))
	jsoAuthen.PutString("txt_password", strings.TrimSpace(jsoHost.GetString("txt_password")))
	jsoCmd.PutObject("jso_authen", jsoAuthen)

	return networks.Request(jsoHost, jsoCmd)
}
