package utils

import (
	"awsvitalcheck/datamodels"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func GetYamlConfigs(path string) datamodels.YamlConfig {
	yConfigs, err := os.ReadFile(path)
	if err != nil {
		log.Println("Error while trying to read the yaml config file, from path: ", path)
		log.Println("Error: ", err)
		os.Exit(1)
	}
	config := datamodels.YamlConfig{}
	err = yaml.Unmarshal(yConfigs, &config)
	if err != nil {
		log.Println("Error unmarshal the yaml to stats object")
		log.Println("Error: ", err)
		os.Exit(1)
	}
	return config
}
