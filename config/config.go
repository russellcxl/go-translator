package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Logger     Logger     `json:"logger"`
	Translator Translator `json:"translator"`
}

type Logger struct {
}

type Translator struct {
	InDir   string `json:"in_dir"`
	OutDir  string `json:"out_dir"`
	OutLang string `json:"out_lang"`
}

func LoadConfig(filePath string) (*Config, error) {
	cfgFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read config file:%s err:%v", filePath, err)
	}

	cfg := &Config{}
	err = json.Unmarshal(cfgFile, cfg)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal config file:%s err:%v", filePath, err)

	}
	return cfg, nil
}
