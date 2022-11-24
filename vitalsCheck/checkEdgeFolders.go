package vitalscheck

import (
	"awsvitalcheck/datamodels"
	"awsvitalcheck/utils"
	"errors"
	"log"
	"os"
	"strings"
)

func CheckEdgeFolders(path string, config *datamodels.YamlConfig) {
	edgeStats := make(map[string]datamodels.EdgeStats)
	log.Println("Opening the edge parent dir: ", path)
	sqHome, err := os.ReadDir(path)
	if err != nil {
		log.Println("Error while trying to open dir: ", path)
		log.Println("Err", err)
		os.Exit(1)
	}

	var stats datamodels.Stats
	var statsArray []map[string]datamodels.EdgeStats
	var resolution string
	for _, edges := range sqHome {
		log.Println("Iterating over the edge dir")
		if edges.IsDir() && !strings.EqualFold(strings.ToLower(edges.Name()), strings.ToLower("vitalStatsCheck")) {
			edgePath := path + "/" + edges.Name()
			log.Println("Opening the edge dir: ", edgePath)
			edgeDir, err := os.ReadDir(edgePath)
			if err != nil {
				log.Println("Error while trying to open dir: ", edgePath)
				log.Println("Err", err)
			}
			log.Println("Getting the last update time of edge folder", path)
			edgetime := utils.GetLastUpdateTime(path, edges.Name())

			var camera []datamodels.EdgeCamera

			log.Println("Iterating over the cameras dir")
			for _, cameras := range edgeDir {
				cameraPath := edgePath + "/" + cameras.Name() + config.LatestImg
				log.Println("Checking the latest jpg in at the following path: ", cameraPath)
				if _, err := os.Stat(cameraPath); errors.Is(err, os.ErrNotExist) {
					log.Println("Latest image is not found in the following path: ", cameraPath)
					log.Println("Error:", err)
					resolution = "latest image does not exist"
				} else {
					log.Println("Checking the resolution of latest.jpg")
					resolution = utils.CheckImageResolution(cameraPath)
				}
				c := datamodels.EdgeCamera{
					Name:       cameras.Name(),
					Resolution: resolution,
				}
				camera = append(camera, c)
				log.Println(">>")
			}
			edgeStats[edges.Name()] = datamodels.EdgeStats{
				Camera:           camera,
				FolderUpdateTime: strings.Split(edgetime, " ")[0],
			}
		}
		log.Println("")
	}
	statsArray = append(statsArray, edgeStats)
	stats.Stats = statsArray
	utils.WriteJsontoStatsFile(stats, config.StatFile)
	log.Println("____________________________________________________________________________________________________________________")
	log.Println("")
}
