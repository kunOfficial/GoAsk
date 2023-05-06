package service

import (
	"GoAsk/config"
	"GoAsk/serializer"
	e "GoAsk/utils/error"
	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	_ "image/png"
	"mime/multipart"
	"os"
	"strconv"
)

type AvatarUploadService struct {
	FileHeader *multipart.FileHeader
	UserID     uint
}

type AvatarDownloadService struct {
	ID int
}

func (service *AvatarUploadService) Upload() serializer.Response {
	exist, _ := DirExist(config.AvatarPath)
	if exist == false {
		err := CreateDir(config.AvatarPath)
		if err != nil {
			return serializer.BuildResponse(e.FileSystemError, nil)
		}
	}
	file, err := service.FileHeader.Open()
	if err != nil {
		return serializer.BuildResponse(e.ErrorOpenImageFile, nil)
	}
	img, _, err := image.Decode(file)
	if err != nil {
		return serializer.BuildResponse(e.ErrorDecodeImageFile, nil)
	}
	resizedImg := resize.Resize(300, 300, img, resize.NearestNeighbor) // 压缩图片文件
	savePath := config.AvatarPath + "user_" + strconv.Itoa(int(service.UserID)) + ".jpeg"
	newFile, err := os.Create(savePath)
	if err != nil {
		return serializer.BuildResponse(e.FileSystemError, nil)
	}
	err = jpeg.Encode(newFile, resizedImg, nil) // 以jpeg文件存储
	if err != nil {
		return serializer.BuildResponse(e.ErrorSaveImageFile, nil)
	}
	return serializer.BuildResponse(e.SUCCESS, nil)
}

func (service *AvatarDownloadService) Download(ctx *gin.Context) serializer.Response {
	filePath := config.AvatarPath + "user_" + strconv.Itoa(int(service.ID)) + ".jpeg"
	ctx.File(filePath)
	return serializer.BuildResponse(e.SUCCESS, nil)
}

// DirExist 判断文件夹是否存在
func DirExist(DirAddr string) (exist bool, err error) {
	s, err := os.Stat(DirAddr)
	if err != nil {
		return false, err
	}
	return s.IsDir(), nil
}

// CreateDir 创建文件夹
func CreateDir(dirName string) (err error) {
	err = os.MkdirAll(dirName, 755) // 递归创建文件夹
	return err
}
