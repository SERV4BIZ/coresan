package api

import (
	"errors"

	"github.com/SERV4BIZ/gfp/jsons"

	"github.com/SERV4BIZ/coresan/api/drivers"
	"github.com/SERV4BIZ/coresan/api/utility"
)

// Unlink is delete file from coresan
func (me *CORESAN) Unlink(txtCSNID string) error {
	me.MutexMapDataItem.RLock()
	dataItem, itemOk := me.MapDataItem[txtCSNID]
	me.MutexMapDataItem.RUnlock()

	if itemOk {
		dataNodeItem := dataItem.DataNode
		jsoReq := drivers.Unlink(dataNodeItem.JSOHost, txtCSNID)
		if jsoReq.GetInt("status") > 0 {
			// Update memory
			me.MutexMapDataItem.Lock()
			delete(me.MapDataItem, txtCSNID)
			me.MutexMapDataItem.Unlock()

			return nil
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

		jsoReq := drivers.Unlink(dataNodeItem.JSOHost, txtCSNID)
		if jsoReq.GetInt("status") > 0 {
			// Update memory
			me.MutexMapDataItem.Lock()
			delete(me.MapDataItem, txtCSNID)
			me.MutexMapDataItem.Unlock()

			return nil
		}
	}

	return errors.New("Not found file")
}
