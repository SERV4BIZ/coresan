package coresans

import (
	"errors"

	"github.com/SERV4BIZ/gfp/jsons"

	"github.com/SERV4BIZ/coresan/api/drivers"
	"github.com/SERV4BIZ/coresan/api/utility"
)

// Info is get info file from coresan
func (me *CORESAN) Info(txtCSNID string) (*jsons.JSONObject, error) {
	me.MutexMapDataItem.RLock()
	dataItem, itemOk := me.MapDataItem[txtCSNID]
	me.MutexMapDataItem.RUnlock()

	if itemOk {
		dataNodeItem := dataItem.DataNode
		jsoReq := drivers.Info(dataNodeItem.JSOHost, txtCSNID)
		if jsoReq.GetInt("status") > 0 {
			jsoReq.GetObject("jso_data").PutString("txt_datanode", dataNodeItem.Name)
			return jsoReq.GetObject("jso_data"), nil
		}

		// Update memory
		me.MutexMapDataItem.Lock()
		delete(me.MapDataItem, txtCSNID)
		me.MutexMapDataItem.Unlock()
	}

	// if not found in dataitem
	me.MutexMapDataNode.RLock()
	jsaNodeKey := jsons.JSONArrayFactory()
	nodeKeys := make([]string, 0)
	for key := range me.MapDataNode {
		jsaNodeKey.PutString(key)
		nodeKeys = append(nodeKeys, key)
	}
	me.MutexMapDataNode.RUnlock()

	for jsaNodeKey.Length() > 0 {
		index := utility.RandomIntn(jsaNodeKey.Length())
		nodeName := jsaNodeKey.GetString(index)
		jsaNodeKey.Remove(index)

		me.MutexMapDataNode.RLock()
		dataNodeItem := me.MapDataNode[nodeName]
		me.MutexMapDataNode.RUnlock()

		jsoReq := drivers.Info(dataNodeItem.JSOHost, txtCSNID)
		if jsoReq.GetInt("status") > 0 {
			// Update memory
			dataItem = new(DataItem)
			dataItem.DataNode = dataNodeItem
			dataItem.CSNID = txtCSNID

			me.MutexMapDataItem.Lock()
			me.MapDataItem[txtCSNID] = dataItem
			me.MutexMapDataItem.Unlock()

			jsoReq.GetObject("jso_data").PutString("txt_datanode", dataNodeItem.Name)
			return jsoReq.GetObject("jso_data"), nil
		}
	}

	return nil, errors.New("Not found file")
}
