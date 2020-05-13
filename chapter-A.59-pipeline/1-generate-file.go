package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/novalagung/gubrak"
)

const totalFile = 3000
const randomStringLength = 5000

var tempPath = filepath.Join(os.Getenv("TEMP"), "chapter-A.59-pipeline-temp")

func main() {
	log.Println("start")
	start := time.Now()

	generate()

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")
}

func generate() {
	os.RemoveAll(tempPath)
	os.MkdirAll(tempPath, os.ModePerm)

	for i := 0; i < totalFile; i++ {
		filename := filepath.Join(tempPath, fmt.Sprintf("file-%d.txt", i))
		content := gubrak.RandomString(randomStringLength)
		err := ioutil.WriteFile(filename, []byte(content), os.ModePerm)
		if err != nil {
			log.Println("Error writing file", filename)
		}

		if i%100 == 0 && i > 0 {
			log.Println(i, "files created")
		}
	}

	log.Printf("%d of total files created", totalFile)
}
