package files

import (
	"net/http"
	"api/app/models"
	"github.com/gin-gonic/gin"
	"strings"
)

func GetAllFilesFromDrive(c *gin.Context) {
	Is.RetrieveAllFilesFromDrive()
}

func GetFileFromDrive(c *gin.Context) {
	fileId := strings.TrimSpace(c.Param("id"))
	urlParams := c.Request.URL.Query()
	word := urlParams["word"][0]

	file, err := Is.SearchInDrive(fileId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "find_error", "description": err.Error()})
		return
	}

	containsWord := strings.Contains(file.Description, word)

	if containsWord {
		c.JSON(200, gin.H{"ok": "word found", "content": file.Description})
	} else {
		c.JSON(404, gin.H{"error": "word not found", "content": file.Description})
	}

	return
}

// Search in local db
func GetFileFromDB(c *gin.Context) {
	fileId := strings.TrimSpace(c.Param("id"))

	if fileId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id_error"})
		return
	}

	file, err := Is.SearchInDB(fileId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "find_error", "description": err.Error()})
		return
	}
	c.JSON(200, file)
	return
}

// Post item in drive ...
func PostFile(c *gin.Context) {
	file := &models.File{}

	if err := c.BindJSON(file); c.Request.ContentLength == 0 || err != nil {
		c.JSON(400, gin.H{"error": "bind_error", "description": err.Error()})
		return
	}

	id, err := Is.CreateFileInDrive(file)

	if err != nil {
		c.JSON(500, gin.H{"error": "post in drive error", "description": err.Error()})
		return
	}

	file.ID = id
	err = Is.CreateFileInDB(file)

	if err != nil {
		c.JSON(500, gin.H{"error": "save_error", "description": err.Error()})
		return
	}

	c.JSON(201, file)
}
