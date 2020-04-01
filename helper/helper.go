package helper

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type (
	PagingData struct {
		Page        int    `json:"page" query:"page" form:"page" validate:"required"`
		Rows        int    `json:"rows" query:"rows" form:"rows" validate:"required"`
		FilterRules string `json:"filterRules" query:"filterRules" form:"filterRules"`
		Sort        string `json:"sort" form:"sort"`
		Order       string `json:"order"`
	}

	FilterRule struct {
		Field string `json:"field" form:"field"`
		Op    string `json:"op" form:"op"`
		Value string `json:"value" form:"value"`
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
