package utils

import (
	"log"
	"os/exec"
)

func GetLastUpdateTime(sqDir, edge string) string {
	s := "ls -l " + sqDir + "| awk '{if (NR!=1 && NF) print $(NF-1),$NF}' | grep " + edge
	log.Println("Running the following command: ", s)
	o, err := exec.Command("bash", "-c", s).Output()
	if err != nil {
		log.Println("Error while trying to run the following commad: ", s)
		log.Println("Error: ", err)
	}
	return string(o)
}

func CheckImageResolution(img string) string {
	s := "file " + img + " | grep -Eo '[[:digit:]]{2,4}+ *x *[[:digit:]]+'"
	log.Println("Running the following command: ", s)
	o, err := exec.Command("bash", "-c", s).Output()
	if err != nil {
		log.Println("Error while trying to run the following commad: ", s)
		log.Println("Error: ", err)
	}
	return string(o)
}
