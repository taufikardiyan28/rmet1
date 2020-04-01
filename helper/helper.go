package helper

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type (
	PagingData struct {
		Page            int    `json:"page" query:"page" form:"page" validate:"required"`
		Show            int    `json:"show" query:"show" form:"show" validate:"required"`
		Search          string `json:"search" query:"search" form:"search"`
		StartCreateDate string `json:"start_create_date" query:"start_create_date" form:"start_create_date"`
		EndCreateDate   string `json:"end_create_date" query:"end_create_date" form:"end_create_date"`
	}

	Search struct {
		Column string `json:"column" form:"column"`
		Op     string `json:"op" form:"op"`
		Value  string `json:"value" form:"value"`
	}

	PeriodFilter struct {
		DateColumn string
		StartDate  string
		EndDate    string
	}

	Configuration struct {
		Database struct {
			Host     string `yaml:"host"`
			Port     int    `yaml:"port"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
			DbName   string `yaml:"db_name"`
		} `yaml:"database"`
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}
)

var Config *Configuration

func ReadConfig() error {
	dat, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		return err
	}

	Config = &Configuration{}
	err = yaml.Unmarshal(dat, Config)

	return err
}
