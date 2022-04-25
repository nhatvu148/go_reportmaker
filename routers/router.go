package routers

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/nhatvu148/go_reportmaker/pkg/upload"
	v1 "github.com/nhatvu148/go_reportmaker/routers/api/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Serve frontend static files
	r.Use(static.Serve("/", static.LocalFile(upload.GetClientBuildFullPath(), true)))

	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	// r.POST("/auth", api.GetAuth)
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// r.POST("/upload", api.UploadImage)

	apiv1 := r.Group("/api/v1")
	// apiv1.Use(jwt.JWT())
	apiv1.Use()
	{
		apiv1.GET("/projects", v1.GetProjects)
		apiv1.POST("/projects", v1.AddProject)
	}

	return r
}
