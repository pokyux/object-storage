package models

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func SaveToLocal(file []byte, ext string) string {
	timestamp := time.Now().UnixNano()
	fileName := fmt.Sprintf("%d.%s", timestamp, ext)
	path := fmt.Sprintf("./files/%s", fileName)
	fileObj, err := os.OpenFile(path, os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer fileObj.Close()

	writer := bufio.NewWriter(fileObj)
	_, err = writer.Write(file)
	if err != nil {
		panic(err)
	}
	writer.Flush()
	return fileName
}
