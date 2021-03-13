package locals

import (
	"fmt"

	"github.com/SERV4BIZ/coresan/server/global"
	"github.com/SERV4BIZ/coresan/server/utility"

	"github.com/SERV4BIZ/gfp/files"
	"github.com/SERV4BIZ/gfp/jsons"
)

// LoadConfig is load json config
func LoadConfig() (*jsons.JSONObject, error) {
	pathfile := fmt.Sprint(utility.GetAppDir(), global.DS, "config.json")
	jsoConfig := jsons.JSONObjectFactory()
	jsoConfig.PutString("txt_host", "localhost")
	jsoConfig.PutInt("int_port", 8765)

	if files.ExistFile(pathfile) {
		var errConfig error
		jsoConfig, errConfig = jsons.JSONObjectFromFile(pathfile)
		if errConfig != nil {
			return nil, errConfig
		}
	}
	return jsoConfig, nil
}
