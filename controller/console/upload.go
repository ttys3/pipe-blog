package console

import (
	"fmt"
	"github.com/b3log/pipe/model"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/b3log/gulu"
	"github.com/b3log/pipe/util"
	"github.com/gin-gonic/gin"
)

func UploadAction(c *gin.Context) {
	result := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, result)

	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		result.Code = util.CodeErr
		result.Msg = fmt.Sprintf("get form err: %s", err.Error())

		return
	}
	files := form.File["file[]"]

	if len(files) == 0 {
		result.Code = util.CodeErr
		result.Msg = "please upload a file"

		return
	}

	var uploadedFiles = make(map[string]string)
	var errFiles = make([]string, 0)
	for _, file := range files {
		bname := filepath.Base(file.Filename)
		filename := strings.ReplaceAll(filepath.Clean(bname), " ", "-")
		fpath := filepath.Join(model.Conf.StoragePath, filename)
		if err := c.SaveUploadedFile(file, fpath); err != nil {
			errFiles = append(errFiles, bname)
		} else {
			uploadedFiles[filename] = "/api/storage/" + filename
		}
	}

	result.Data = map[string]interface{}{
		"succMap": uploadedFiles,
		"errFiles": errFiles,
	}
	result.Code = util.CodeOk
	result.Msg =  fmt.Sprintf("uploaded successfully %d files", len(files))
}
