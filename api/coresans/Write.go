package api

import (
	"errors"
	"fmt"

	"github.com/SERV4BIZ/gfp/jsons"

	"github.com/SERV4BIZ/coresan/api/drivers"
	"github.com/SERV4BIZ/coresan/api/utility"
)

// Write is write file to coresan
func (me *CORESAN) Write(txtFilename string, dblExpire float64, buffer []byte) (*jsons.JSONObject, error) {
	me.MutexMapDataNode.RLock()
	jsaNodeKey := jsons.JSONArrayFactory()
	nodeKeys := make([]string, 0)
	for key := range me.MapDataNode {
		jsaNodeKey.PutString(key)
		nodeKeys = append(nodeKeys, key)
	}
	me.MutexMapDataNode.RUnlock()

	// If not found then insert row
	me.MutexMapDataNode.RLock()
	dataNodeItem := me.MapDataNode[nodeKeys[utility.RandomIntn(len(nodeKeys))]]
	me.MutexMapDataNode.RUnlock()

	jsoReq := drivers.Write(dataNodeItem.JSOHost, txtFilename, dblExpire, buffer)
	if jsoReq.GetInt("status") == 0 {
		return nil, errors.New(fmt.Sprint("Can not write to coresan in node ", dataNodeItem.Name, " [ ", jsoReq.GetString("txt_msg"), " ]"))
	}

	// Update memory
	dataItem := new(DataItem)
	dataItem.DataNode = dataNodeItem
	dataItem.CSNID = jsoReq.GetObject("jso_data").GetString("txt_csnid")

	me.MutexMapDataItem.Lock()
	me.MapDataItem[dataItem.CSNID] = dataItem
	me.MutexMapDataItem.Unlock()

	return jsoReq.GetObject("jso_data"), nil
}
