package models

import (
	"github.com/jinzhu/gorm"
)

type Project struct {
	Pjid           string `json:"pjid"`
	Pjname_jp      string `json:"pjname_jp"`
	Pjname_en      string `json:"pjname_en"`
	Startdate      string `json:"startdate"`
	Deadline       string `json:"deadline"`
	Expecteddate   string `json:"expecteddate"`
	Scode          int    `json:"scode"`
	Ccode          int    `json:"ccode"`
	Cpjname_jp     string `json:"cpjname_jp"`
	Cpjname_en     string `json:"cpjname_en"`
	Rcode          string `json:"rcode"`
	Completiondate string `json:"completiondate"`
}

func (Project) TableName() string {
	return "t_projectmaster"
}

func GetProjects() ([]*Project, error) {
	var projects []*Project
	// SELECT pjid, pjname_en, pjname_jp FROM projectdata.t_projectmaster WHERE scode = 0
	err := db.Select("pjid, pjname_en, pjname_jp").Where("scode = ?", 0).Find(&projects).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return projects, nil
}
