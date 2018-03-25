package files

import (
	"net/http"
	"api/app/models"
	"github.com/gin-gonic/gin"
	"strings"
)

// Search word in drive ...
func Search(c *gin.Context) {
	c.JSON(200, nil)
	return
}

// Search in local db
func Get(c *gin.Context){
	fileId := strings.TrimSpace(c.Param("id"))

	if fileId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id_error"})
		return
	}

	file, err := Is.GetFile(fileId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "find_error", "description": err.Error()})
		return
	}
	c.JSON(200, file)
	return
}

// Post item in drive ...
func Create(c *gin.Context) {
	f := &models.File{}
	if err := c.BindJSON(f); c.Request.ContentLength == 0 || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bind_error", "description": err.Error()})
		return
	}
	err := Is.CreateFile(f)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "save_error", "description": err.Error()})
		return
	}
	c.JSON(201, f)
}
