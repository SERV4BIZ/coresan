package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/SERV4BIZ/coresan/api/coresans"
	"github.com/SERV4BIZ/gfp/jsons"
)

func GetAppDir() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath, _ := filepath.Abs(filepath.Dir(ex))
	return exPath
}

func main() {
	pathFile := fmt.Sprint(GetAppDir(), "/config.json")
	jsoConfig, _ := jsons.JSONObjectFromFile(pathFile)
	csnItem, _ := coresans.Factory(jsoConfig)

	jsoWrite, errWrite := csnItem.Write("test.json", 0, []byte(jsoConfig.ToString()))
	if errWrite != nil {
		panic(errWrite)
	}

	txtCSNID := jsoWrite.GetString("txt_csnid")

	fmt.Println("Write ===")
	fmt.Println(jsoWrite.ToString())

	jsoRewrite, errRewrite := csnItem.Rewrite(txtCSNID, "test.json", 0, []byte("i love you"))
	if errRewrite != nil {
		panic(errRewrite)
	}

	fmt.Println("Rwrite ===")
	fmt.Println(jsoRewrite.ToString())

	fmt.Println("Info ===")
	jsoInfo, _ := csnItem.Info(txtCSNID)
	fmt.Println(jsoInfo.ToString())

	fmt.Println("Read ===")
	jsoRead, buffer, _ := csnItem.Read(txtCSNID)
	fmt.Println(jsoRead.ToString())
	fmt.Println(string(buffer))

	/*fmt.Println("Exist ===")
	errExist := csnItem.Exist(txtCSNID)
	fmt.Println(errExist)

	fmt.Println("Unlink ===")
	errUnlink := csnItem.Unlink(txtCSNID)
	fmt.Println(errUnlink)

	fmt.Println("Exist Again ===")
	errExist = csnItem.Exist(txtCSNID)
	fmt.Println(errExist)

	fmt.Println("Unlink ===")
	errUnlink = csnItem.Unlink(txtCSNID)
	fmt.Println(errUnlink)*/
}
