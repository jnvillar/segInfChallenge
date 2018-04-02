package files

import (
	"api/app/models"
	"database/sql"
	"github.com/gin-gonic/gin"
)

var (
	Is models.FileServiceInterface
)

func Configure(r *gin.Engine, db *sql.DB) {
	Is = &FilesService{DB: db}

	r.GET("/search-in-doc/:id", GetFileFromDrive)
	r.GET("/showAll", GetAllFilesFromDrive)
	r.GET("/file/:id", GetFileFromDB)
	r.POST("/file", PostFile)
}
