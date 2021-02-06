package locals

import (
	"fmt"
	"strings"

	"github.com/SERV4BIZ/coresan/config/utility"
	"github.com/SERV4BIZ/gfp/files"
	"github.com/SERV4BIZ/gfp/jsons"
)

// ListDataNode is listing data node
func ListDataNode() (*jsons.JSONArray, error) {
	jsaList := jsons.JSONArrayFactory()
	pathdir := fmt.Sprint(utility.GetAppDir(), utility.DS, "datanodes")
	filelist, err := files.ScanDir(pathdir)
	if err != nil {
		return nil, err
	}

	for _, fileName := range filelist {
		if strings.HasSuffix(fileName, ".json") {
			jsaList.PutString(strings.TrimSpace(strings.ToLower(strings.TrimSuffix(fileName, ".json"))))
		}
	}
	return jsaList, err
}
