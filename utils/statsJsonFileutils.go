package utils

import (
	"awsvitalcheck/datamodels"
	"encoding/json"
	"log"
	"os"
)

func WriteJsontoStatsFile(stats datamodels.Stats, statsFile string) {
	log.Println("Marshalling the object to json string")
	jsonStr, err := json.Marshal(stats)
	if err != nil {
		log.Println("Error while trying to convert object to json")
		log.Println("Error: ", err)
		os.Exit(1)
	} else {
		log.Println("Writing json data to 'stats.json'")
		err = os.WriteFile(statsFile, jsonStr, 0644)
		if err != nil {
			log.Println("Error while trying to create stats.json file")
			log.Println("Error: ", err)
			os.Exit(1)
		}
	}
}
