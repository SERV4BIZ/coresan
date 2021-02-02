package drivers

import (
	"strings"

	"github.com/SERV4BIZ/coresan/api/networks"
	"github.com/SERV4BIZ/gfp/jsons"
)

// Unlink is delete file from coresan
func Unlink(jsoHost *jsons.JSONObject, txtCSNID string) *jsons.JSONObject {
	jsoCmd := jsons.JSONObjectFactory()
	jsoCmd.PutString("txt_command", "storage_unlink")
	jsoCmd.PutString("txt_csnid", strings.ToLower(strings.TrimSpace(txtCSNID)))

	jsoAuthen := jsons.JSONObjectFactory()
	jsoAuthen.PutString("txt_username", strings.TrimSpace(strings.ToLower(jsoHost.GetString("txt_username"))))
	jsoAuthen.PutString("txt_password", strings.TrimSpace(jsoHost.GetString("txt_password")))
	jsoCmd.PutObject("jso_authen", jsoAuthen)

	return networks.Request(jsoHost, jsoCmd)
}
