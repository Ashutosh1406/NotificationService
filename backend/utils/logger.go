package utils

import (
	"log"
	"os"
)

var Logger = log.New(os.Stdout, "Info:", log.Ldate|log.Ltime|log.Lshortfile)

func logError(err error) {
	if err != nil {
		Logger.Println("ERROR:", err)
	}
}
