package utils

import (
	"log"
	"os"
)

//获取当前项目的根目录
func GetBaseDir() string {
	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatalf("can't get the current work path")
	}
	return currentPath
}
