package saveFile

import (
	// "fmt"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/cn-joyconn/goadmin/models/global"
	
	snowflake "github.com/cn-joyconn/goutils/snowflake"
)

func GetSaveFilePath(file *multipart.FileHeader, cfg *global.UploadCfg) (string, string) {
	// newID := global.SnowflakeWorker.GetId()
	newID := snowflake.NextId()
	fullFileName := file.Filename
	suffixName := fullFileName[strings.LastIndex(fullFileName, "."):]
	// newFileName := strconv.FormatInt(newID, 16) + suffixName
	newFileName := strconv.FormatUint(newID, 16) + suffixName
	// today := time.Now().Format("yyyyMMdd")
	basicPath := "/" + time.Now().Format("yyyyMMdd") + "/"
	newUploadPath := cfg.SavePath + basicPath
	newUploadPath = strings.ReplaceAll(newUploadPath, "//", "/")
	os.MkdirAll(path.Dir(newUploadPath), os.ModePerm)
	newFilePath := newUploadPath + newFileName
	returnUrl := cfg.VisitDomain + cfg.VisitPath + basicPath + "/" + newFileName
	returnUrl = strings.ReplaceAll(returnUrl, "//", "/")
	return newFilePath, returnUrl
	// os.MkdirAll(path.Dir(newUploadPath), os.ModePerm)
	// out, err := os.Create(newFilePath)
	// defer out.Close()

}
