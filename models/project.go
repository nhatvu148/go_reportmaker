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

func AddProject(data map[string]interface{}) error {
	project := Project{
		Pjid:           data["pjid"].(string),
		Pjname_jp:      data["pjname_jp"].(string),
		Pjname_en:      data["pjname_en"].(string),
		Startdate:      data["startdate"].(string),
		Deadline:       data["deadline"].(string),
		Expecteddate:   data["expecteddate"].(string),
		Scode:          data["scode"].(int),
		Ccode:          data["ccode"].(int),
		Cpjname_jp:     data["cpjname_jp"].(string),
		Cpjname_en:     data["cpjname_en"].(string),
		Rcode:          data["rcode"].(string),
		Completiondate: data["completiondate"].(string),
	}
	if err := db.Create(&project).Error; err != nil {
		return err
	}

	return nil
}
