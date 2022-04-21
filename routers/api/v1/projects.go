package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nhatvu148/go_reportmaker/pkg/app"
	"github.com/nhatvu148/go_reportmaker/pkg/e"
	"github.com/nhatvu148/go_reportmaker/service/project_service"
)

func GetProjects(c *gin.Context) {
	appG := app.Gin{C: c}

	projectService := project_service.Project{}

	project, err := projectService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, project)

	// appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
	// 	"message":  "pong",
	// 	"message2": "pong2",
	// })
}
