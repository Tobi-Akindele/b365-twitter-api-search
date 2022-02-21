package config

import (
	"bet365/utils"
	"bufio"
	"log"
	"os"
	"strings"
)

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		utils.CheckError(err)
	}
}

func Loader(envFileName string) (err error) {
	file, err := os.Open(envFileName)
	utils.CheckError(err)

	defer closeFile(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineContent := strings.Trim(scanner.Text(), " ")
		if !strings.HasPrefix(lineContent, "#") && len(lineContent) > 0 && strings.Contains(lineContent, "=") {
			envSeperator := strings.Index(lineContent, "=")
			key := lineContent[:envSeperator]
			value := lineContent[envSeperator+1:]
			_ = os.Setenv(key, value)
		}
	}
	log.Println("Environment variables loaded successfully")
	return
}
