package main

import (
	"awsvitalcheck/utils"
	vitalscheck "awsvitalcheck/vitalsCheck"
	"log"
	"os"
)

func main() {
	currentDir := os.Args[1]
	configs := utils.GetYamlConfigs(currentDir)
	f, err := os.OpenFile(configs.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println("Starting vital check")
	vitalscheck.CheckEdgeFolders(configs.EdgesParent, &configs)
}
