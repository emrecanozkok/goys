package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"main/logger"
	"os"
	"time"
)

var (
	GlobalStore = make(map[string]string)
)

const fileName = "%d_backup.json"
const fileFolder = "./tmp/"

func FindLatestFile() string {

	files, _ := ioutil.ReadDir(fileFolder)
	var newestFile string = ""
	var newestTime int64 = 0
	for _, f := range files {
		fi, err := os.Stat(fileFolder + f.Name())
		if err != nil {
			logger.Log.Printf(err.Error())
			return ""
		}
		currTime := fi.ModTime().Unix()
		if currTime > newestTime {
			newestTime = currTime
			newestFile = f.Name()
		}
	}

	return newestFile
}

func LoadFromFile() bool {
	latestFile := FindLatestFile()
	if latestFile == "" {
		return false
	}

	fileFullPath := fmt.Sprintf("%s%s", fileFolder, latestFile)
	jsonFile, err := os.ReadFile(fileFullPath)
	// if we os.Open returns an error then handle it
	if err != nil {
		logger.Log.Printf(err.Error())
		return false
	}
	err = json.Unmarshal(jsonFile, &GlobalStore)
	if err != nil {
		logger.Log.Println("error during loading json file")
		return false
	}
	logger.Log.Printf("successfully Opened %s file", fileFullPath)
	return true
}
func FileExists(fileFullPath string) bool {
	_, err := os.Stat(fileFullPath)
	if err == nil {
		return true
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}

func DeleteFile(fileFullPath string) {
	e := os.Remove(fileFullPath)
	if e != nil {
		logger.Log.Printf(e.Error())
	}
}

func DumpDataToFile() bool {
	now := time.Now()
	jsonData, err := json.Marshal(GlobalStore)
	fileFullPath := fmt.Sprintf(fileFolder+fileName, now.Unix())
	if err != nil {
		panic(err)
	}
	if FileExists(fileFullPath) == true {
		logger.Log.Printf("file already exist, deleting")
		DeleteFile(fileFullPath)
	}
	jsonFile, err := os.Create(fileFullPath)

	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	jsonFile.Write(jsonData)
	logger.Log.Printf("content wrote to %s", fileFullPath)
	jsonFile.Close()

	return true
}
