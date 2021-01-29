package locals

import (
	"errors"
	"fmt"
	"strings"

	"github.com/SERV4BIZ/coresan/config/utility"
	"github.com/SERV4BIZ/gfp/filesystem"
	"github.com/SERV4BIZ/gfp/jsons"
)

// LoadDataNodeInfo is load data node info
func LoadDataNodeInfo(name string) (*jsons.JSONObject, error) {
	pathfile := fmt.Sprint(utility.GetAppDir(), utility.DS, "datanodes", utility.DS, strings.TrimSpace(strings.ToLower(name)), ".json")
	if filesystem.ExistFile(pathfile) {
		return jsons.JSONObjectFromFile(pathfile)
	}
	return nil, errors.New(fmt.Sprint("Not found ", pathfile, " file"))
}
