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

func (a *Project) Add() error {
	project := map[string]interface{}{
		"pjid":           a.Pjid,
		"pjname_jp":      a.Pjname_jp,
		"pjname_en":      a.Pjname_en,
		"startdate":      a.Startdate,
		"deadline":       a.Deadline,
		"expecteddate":   a.Expecteddate,
		"scode":          a.Scode,
		"ccode":          a.Ccode,
		"cpjname_jp":     a.Cpjname_jp,
		"cpjname_en":     a.Cpjname_en,
		"rcode":          a.Rcode,
		"completiondate": a.Completiondate,
	}

	if err := models.AddProject(project); err != nil {
		return err
	}

	return nil
}

func (a *Project) GetAll() ([]*models.Project, error) {
	var projects []*models.Project

	projects, err := models.GetProjects()
	if err != nil {
		return nil, err
	}

	return projects, nil
}
