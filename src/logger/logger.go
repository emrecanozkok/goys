package logger

import (
	"fmt"
	"log"
	"main/config"
	"os"
)

var (
	Log *log.Logger
)
func init() {
	// set location of log file
	var logpath = fmt.Sprintf("%s%s",config.LOG_PATH,"/go-ys.log")

	var file, err1 = os.Create(logpath)


	if err1 != nil {
		panic(err1)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	Log.Println("LogFile : " + logpath)
}
