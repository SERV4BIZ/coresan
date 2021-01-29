package global

import (
	"fmt"
	"strings"
	"sync"

	"github.com/SERV4BIZ/coresan/config/utility"
	"github.com/SERV4BIZ/gfp/files"
	"github.com/SERV4BIZ/gfp/jsons"
)

// AppName is name of application
var AppName string = "CORESAN SERVER"

// AppVersion is version of application
var AppVersion string = "1.0.0"

// DS is split of path
var DS string = "/"

// MutexJSOConfig is mutex lock of JSOConfig
var MutexJSOConfig sync.RWMutex

// JSOConfig is config json object
var JSOConfig *jsons.JSONObject

// MutexState is mutex lock of MemoryState
var MutexState sync.RWMutex

// MemoryState is state of memory
var MemoryState int = 0

// LoadState is load of request
var LoadState int = 0

// CountState is count of request per second
var CountState int = 0

// Username is authen username
var Username = ""

// Password is authen password
var Password = ""

// NFSPath is base nfs path
var NFSPath = ""

// LoadConfig is load json config
func LoadConfig() (*jsons.JSONObject, error) {
	pathfile := fmt.Sprint(utility.GetAppDir(), DS, "config.json")
	jsoConfig := jsons.JSONObjectFactory()
	jsoConfig.PutString("txt_host", "localhost")
	jsoConfig.PutInt("num_port", 5679)

	if files.ExistFile(pathfile) {
		var errConfig error
		jsoConfig, errConfig = jsons.JSONObjectFromFile(pathfile)
		if errConfig != nil {
			return nil, errConfig
		}
	}
	return jsoConfig, nil
}

// GetJSOConfig is get copy json object
func GetJSOConfig() (*jsons.JSONObject, error) {
	MutexJSOConfig.Lock()
	defer MutexJSOConfig.Unlock()
	return JSOConfig.Copy()
}

// GetFullPath is get fullpath
func GetFullPath(txtUUID string) string {
	txtPath := strings.ReplaceAll(txtUUID, "-", DS)
	return fmt.Sprint(NFSPath, DS, txtPath)
}
