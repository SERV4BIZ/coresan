package api

import (
	"sync"

	"github.com/SERV4BIZ/gfp/jsons"
)

// DataNode is struct for datanode info
type DataNode struct {
	sync.RWMutex
	CORESAN *CORESAN

	Name    string
	JSOHost *jsons.JSONObject
}

// DataItem is struct in Cabinet object
type DataItem struct {
	DataNode *DataNode
	CSNID    string
}

// CORESAN is main object
type CORESAN struct {
	sync.RWMutex

	UUID          string
	JSOConfigHost *jsons.JSONObject

	MutexMapDataNode sync.RWMutex
	MapDataNode      map[string]*DataNode

	MutexMapDataItem sync.RWMutex
	MapDataItem      map[string]*DataItem
}
