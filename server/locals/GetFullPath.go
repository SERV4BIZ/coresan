package locals

import (
	"fmt"
	"strings"

	"github.com/SERV4BIZ/coresan/server/global"
)

// GetFullPath is get fullpath
func GetFullPath(txtUUID string) string {
	txtPath := strings.ReplaceAll(txtUUID, "-", global.DS)
	return fmt.Sprint(global.NFSPath, global.DS, txtPath)
}
