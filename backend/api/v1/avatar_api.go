package v1

import (
	"GoAsk/serializer"
	"GoAsk/service"
	e "GoAsk/utils/error"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// UploadAvatar 上传头像
func UploadAvatar(c *gin.Context) {
	var avatarUploadService service.AvatarUploadService
	userID, exists := c.Get("user_id")
	if exists == false {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}
	if fileHeader, err := c.FormFile("avatar"); err != nil {
		c.JSON(400, serializer.BuildResponse(e.ErrorOpenImageFile, nil))
	} else {
		avatarUploadService.FileHeader = fileHeader
		avatarUploadService.UserID = userID.(uint)
		res := avatarUploadService.Upload()
		c.JSON(http.StatusOK, res)
	}

}

// GetAvatar 下载头像（其实没啥用）
func GetAvatar(c *gin.Context) {
	var avatarDownloadService service.AvatarDownloadService
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, serializer.BuildResponse(e.InvalidParams, nil))
	} else {
		avatarDownloadService.ID, _ = strconv.Atoi(id)
		res := avatarDownloadService.Download(c)
		c.JSON(http.StatusOK, res)
	}
}
