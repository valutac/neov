package cfg

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/Valutac/neov"
)

type Configuration struct {
	Credential struct {
		Host        string `toml:"host"`
		Username    string `toml:"username"`
		Password    string `toml:"password"`
		ProjectID   string `toml:"project_id"`
		ProjectName string `toml:"project_name"`
	} `toml:"credential"`
}

func Init(path string) *Configuration {
	var config Configuration
	if _, err := toml.DecodeFile(path, &config); err != nil {
		log.Fatalf("[FATAL] error opening configuration file: %s\n", err.Error())
	}
	if config.Credential.Host == "" {
		config.Credential.Host = neov.AuthURL
	}
	return &config
}
