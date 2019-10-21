package controller

import (
	"github.com/b3log/pipe/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

// UpdateThemeAction updates theme.
func storageAction(c *gin.Context) {

	filename := c.Param("filename")
	filename = filepath.Clean(filename)

	fpath := filepath.Join(model.Conf.StoragePath, filename)
	if _, err := os.Stat(fpath); err == nil {
		c.File(fpath)
	} else {
		c.String(http.StatusNotFound, "file not found: %s",err)
	}
}