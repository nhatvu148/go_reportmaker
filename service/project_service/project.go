package project_service

import (
	"github.com/nhatvu148/go_reportmaker/models"
)

type Project struct {
	Pjid           string
	Pjname_jp      string
	Pjname_en      string
	Startdate      string
	Deadline       string
	Expecteddate   string
	Scode          int
	Ccode          int
	Cpjname_jp     string
	Cpjname_en     string
	Rcode          string
	Completiondate string
}

func (a *Project) GetAll() ([]*models.Project, error) {
	var projects []*models.Project

	projects, err := models.GetProjects()
	if err != nil {
		return nil, err
	}

	return projects, nil
}
