package global

import (
	"sync"
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

// Username is authen username
var Username = ""

// Password is authen password
var Password = ""

// NFSPath is base nfs path
var NFSPath = ""

// MaxRead is Max MultiPart of body in request
var MaxRead int = 0

// MutexState is mutex lock of MemoryState
var MutexState sync.RWMutex

// MemoryState is state of memory
var MemoryState int = 0

// LoadState is load of request
var LoadState int = 0

// CountState is count of request per second
var CountState int = 0
