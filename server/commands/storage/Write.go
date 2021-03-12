package storage

import (
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/SERV4BIZ/coresan/server/global"
	"github.com/SERV4BIZ/gfp/datetimes"
	"github.com/SERV4BIZ/gfp/files"
	"github.com/SERV4BIZ/gfp/jsons"
	"github.com/SERV4BIZ/gfp/uuid"
)

// Write is write data to coresan
func Write(jsoCmd *jsons.JSONObject) *jsons.JSONObject {
	jsoResult := jsons.JSONObjectFactory()
	jsoResult.PutInt("status", 0)

	txtUUID, errUUID := uuid.NewV4()
	if errUUID != nil {
		jsoResult.PutString("txt_msg", fmt.Sprint("Can not generate uuid [ ", errUUID, " ]"))
		return jsoResult
	}
	txtUUID = strings.ReplaceAll(txtUUID, "-", "")

	dateTime := datetimes.DateTimeFactory()
	dateTime.UTC()
	txtCSNID := fmt.Sprintf("%04d%02d%02d-%02d%02d-%s", dateTime.GetYear(), dateTime.GetMonth(), dateTime.GetDay(), dateTime.GetHour(), dateTime.GetMinute(), txtUUID)

	txtFilename := jsoCmd.GetString("txt_filename")
	dblExpire := jsoCmd.GetDouble("dbl_expire")
	txtData := jsoCmd.GetString("txt_data")
	txtExt := "dat"
	exts := strings.Split(txtFilename, ".")
	if len(exts) >= 2 {
		txtExt = strings.ToLower(exts[len(exts)-1])
	}

	txtFullpath := global.GetFullPath(txtCSNID)
	txtInfopath := fmt.Sprint(txtFullpath, global.DS, "info.json")
	txtDatapath := fmt.Sprint(txtFullpath, global.DS, "data.dat")

	buffer, errDecode := base64.StdEncoding.DecodeString(txtData)
	if errDecode != nil {
		jsoResult.PutString("txt_msg", fmt.Sprint("Can not decode base64 from data string [ ", errDecode, " ]"))
		return jsoResult
	}

	errMake := files.MakeDir(txtFullpath)
	if errMake != nil {
		jsoResult.PutString("txt_msg", fmt.Sprint("Can not make directory [ ", errMake, " ]"))
		return jsoResult
	}

	intSize, errWrite := files.WriteFile(txtDatapath, buffer)
	if errWrite != nil {
		jsoResult.PutString("txt_msg", fmt.Sprint("Can not write data [ ", errWrite, " ]"))
		return jsoResult
	}

	jsoInfo := jsons.JSONObjectFactory()
	jsoInfo.PutString("txt_csnid", txtCSNID)
	jsoInfo.PutString("txt_name", txtFilename)
	jsoInfo.PutString("txt_ext", txtExt)
	jsoInfo.PutDouble("dbl_stamp", float64(time.Now().Unix()))
	jsoInfo.PutDouble("dbl_expire", dblExpire)
	jsoInfo.PutInt("int_size", intSize)
	intSize, errToFile := jsoInfo.ToFile(txtInfopath)
	if errToFile != nil {
		jsoResult.PutString("txt_msg", fmt.Sprint("Can not export json to file [ ", errToFile, " ]"))
		return jsoResult
	}

	jsoResult.PutObject("jso_data", jsoInfo)
	jsoResult.PutInt("status", 1)
	return jsoResult
}
