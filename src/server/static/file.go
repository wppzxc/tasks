package static

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/wppzxc/tasks/src/database"
	"github.com/wppzxc/tasks/src/pkg/errors"
	"github.com/wppzxc/tasks/src/pkg/types"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
)

const (
	fileServer     = "http://212.64.87.242"
	serverPath     = "/static/"
	staticFilePath = "/data/static/"
	imgJPEG        = ".jpeg"
	imgJPG         = ".jpg"
	imgPNG         = ".png"
)

func GetServiceInfo(ctx echo.Context) error {
	svc, err := GetCurrentServiceInfo()
	if err != nil {
		return ctx.JSON(errors.NewTaskServerError(http.StatusNotFound, err))
	}
	return ctx.JSON(http.StatusOK, svc)
}

func UpdateServiceInfo(ctx echo.Context) error {
	username := ctx.Param("username")
	if len(username) == 0 || username != types.AdminUser {
		return ctx.JSON(http.StatusNotFound, "not admin user")
	}
	svc := new(database.ServiceInfo)
	if err := ctx.Bind(svc); err != nil {
		return ctx.JSON(errors.NewTaskServerError(http.StatusInternalServerError, err))
	}
	newSvc, err := UpdateService(svc)
	if err != nil {
		return ctx.JSON(errors.NewTaskServerError(http.StatusInternalServerError, err))
	}
	return ctx.JSON(http.StatusOK, newSvc)
}

func GetCurrentServiceInfo() (*database.ServiceInfo, error) {
	svc, err := database.GetCurrentServiceInfo()
	if err != nil {
		return nil, err
	}
	return svc, nil
}

func UpdateService(svc *database.ServiceInfo) (*database.ServiceInfo, error) {
	svc, err := database.UpdateServiceInfo(svc)
	if err != nil {
		return nil, err
	}
	return svc, nil
}

func UploadFiles(ctx echo.Context) error {
	username := ctx.Param("username")
	if len(username) == 0 {
		return ctx.JSON(errors.NewTaskServerError(http.StatusNotFound, fmt.Errorf("username must provide")))
	}

	files := getFileList(ctx)

	saveFunc := func(file *multipart.FileHeader) string {
		if file == nil {
			return ""
		}
		ext := path.Ext(file.Filename)
		if ext != imgJPEG && ext != imgJPG && ext != imgPNG {
			return ""
		}

		data, err := file.Open()
		if err != nil {
			return ""
		}
		defer data.Close()

		now := strconv.FormatInt(time.Now().UnixNano(), 10)
		filename := username + "-" + now + ext
		savePath := staticFilePath + filename
		saveFile, err := os.Create(savePath)
		if err != nil {
			return ""
		}
		defer saveFile.Close()

		if _, err := io.Copy(saveFile, data); err != nil {
			return ""
		}
		fmt.Println("save img : ", savePath)
		return fileServer + serverPath + filename
	}

	var result []string
	for _, file := range files {
		url := saveFunc(file)
		result = append(result, url)
	}
	return ctx.JSON(http.StatusOK, result)
}

func getFileList(ctx echo.Context) []*multipart.FileHeader {
	var result []*multipart.FileHeader
	for i := 0; i < 10; i++ {
		index := "img" + strconv.Itoa(i)
		file, err := ctx.FormFile(index)
		if err != nil {
			break
		} else {
			result = append(result, file)
		}
	}
	return result
}
