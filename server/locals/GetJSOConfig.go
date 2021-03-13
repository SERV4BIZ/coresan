package locals

import (
	"github.com/SERV4BIZ/coresan/server/global"

	"github.com/SERV4BIZ/gfp/jsons"
)

// GetJSOConfig is get copy json object
func GetJSOConfig() (*jsons.JSONObject, error) {
	global.MutexJSOConfig.Lock()
	defer global.MutexJSOConfig.Unlock()
	return global.JSOConfig.Copy()
}
