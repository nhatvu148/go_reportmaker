package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nhatvu148/go_reportmaker/pkg/app"
	"github.com/nhatvu148/go_reportmaker/pkg/e"
	"github.com/nhatvu148/go_reportmaker/service/project_service"
)

type AddProjectForm struct {
	Pjid           string `form:"pjid" valid:"Required"`
	Pjname_jp      string `form:"pjname_jp" valid:"Required"`
	Pjname_en      string `form:"pjname_en" valid:"Required"`
	Startdate      string `form:"startdate" valid:"Required"`
	Deadline       string `form:"deadline" valid:"Required"`
	Expecteddate   string `form:"expecteddate" valid:"Required"`
	Scode          int    `form:"scode" valid:"Required"`
	Ccode          int    `form:"ccode" valid:"Required"`
	Cpjname_jp     string `form:"cpjname_jp" valid:"Required"`
	Cpjname_en     string `form:"cpjname_en" valid:"Required"`
	Rcode          string `form:"rcode" valid:"Required"`
	Completiondate string `form:"completiondate" valid:"Required"`
}

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

func AddProject(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddProjectForm
	)

	if err := c.ShouldBindJSON(&form); err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	// httpCode, errCode := app.BindAndValid(c, &form)
	// if errCode != e.SUCCESS {
	// 	appG.Response(httpCode, errCode, nil)
	// 	return
	// }

	projectService := project_service.Project{
		Pjid:           form.Pjid,
		Pjname_jp:      form.Pjname_jp,
		Pjname_en:      form.Pjname_en,
		Startdate:      form.Startdate,
		Deadline:       form.Deadline,
		Expecteddate:   form.Expecteddate,
		Scode:          form.Scode,
		Ccode:          form.Ccode,
		Cpjname_jp:     form.Cpjname_jp,
		Cpjname_en:     form.Cpjname_en,
		Rcode:          form.Rcode,
		Completiondate: form.Completiondate,
	}
	if err := projectService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
