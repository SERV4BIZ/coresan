package coresans

import (
	"errors"
	"fmt"
	"strings"

	"github.com/SERV4BIZ/coresan/api/networks"
	"github.com/SERV4BIZ/gfp/jsons"
	"github.com/SERV4BIZ/gfp/uuid"
)

// Factory is begin HScaleDB object
func Factory(jsoConfigHost *jsons.JSONObject) (*CORESAN, error) {
	myUUID, errUUID := uuid.NewV4()
	if errUUID != nil {
		return nil, errUUID
	}

	csnItem := new(CORESAN)
	csnItem.UUID = myUUID
	csnItem.JSOConfigHost = jsoConfigHost
	csnItem.MapDataNode = make(map[string]*DataNode)
	csnItem.MapDataItem = make(map[string]*DataItem)
	fmt.Println("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * *")
	fmt.Println("CoreSAN Cluster Factory")
	fmt.Println("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * *")
	fmt.Println(fmt.Sprint("UUID : ", csnItem.UUID))
	fmt.Println("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * *")
	fmt.Println("Loading datanode info.")
	fmt.Println("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * *")

	// Get Data node info
	jsaDataNode, errDataNode := csnItem.DataNodeListing()
	if errDataNode != nil {
		return nil, errDataNode
	}

	for i := 0; i < jsaDataNode.Length(); i++ {
		buff := fmt.Sprint(i+1, " ) ", jsaDataNode.GetString(i))
		fmt.Println(buff)

		jsoNodeInfo, errNodeInfo := csnItem.DataNodeInfo(jsaDataNode.GetString(i))
		if errNodeInfo != nil {
			return nil, errNodeInfo
		}

		nNodeItem := new(DataNode)
		nNodeItem.CORESAN = csnItem
		nNodeItem.Name = strings.ToLower(strings.TrimSpace(jsoNodeInfo.GetString("txt_name")))
		nNodeItem.JSOHost = jsoNodeInfo.GetObject("jso_coresan")

		jsoReq := networks.Ping(nNodeItem.JSOHost)
		if jsoReq.GetInt("status") == 0 {
			return nil, errors.New("Can not ping network")
		}

		csnItem.MutexMapDataNode.Lock()
		csnItem.MapDataNode[nNodeItem.Name] = nNodeItem
		csnItem.MutexMapDataNode.Unlock()
	}
	fmt.Println("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * *")

	return csnItem, nil
}
