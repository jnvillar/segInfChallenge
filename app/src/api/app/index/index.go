package index

import (
	"github.com/gin-gonic/gin"
)

// Configure for index
func Configure(r *gin.Engine) {
	r.GET("/", index)
}
