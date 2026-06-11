package main

import (
	"crypto/md5"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var tempPath = filepath.Join(os.Getenv("TEMP"), "chapter-A.62-concurrency-pipeline")

type FileInfo struct {
	FilePath  string
	Content   []byte
	Sum       string
	IsRenamed bool
}

func main() {
	log.Println("start")
	start := time.Now()

	chanFileContent := readFilesV2()

	chanFileSum1 := getSum(chanFileContent)
	chanFileSum2 := getSum(chanFileContent)
	chanFileSum3 := getSum(chanFileContent)
	chanFileSum := mergeChanFileInfo(chanFileSum1, chanFileSum2, chanFileSum3)

	chanRename1 := rename(chanFileSum)
	chanRename2 := rename(chanFileSum)
	chanRename3 := rename(chanFileSum)
	chanRename4 := rename(chanFileSum)
	chanRename := mergeChanFileInfo(chanRename1, chanRename2, chanRename3, chanRename4)

	counterRenamed := 0
	counterTotal := 0
	for fileInfo := range chanRename {
		if fileInfo.IsRenamed {
			counterRenamed++
		}
		counterTotal++
	}

	log.Printf("%d/%d files renamed", counterRenamed, counterTotal)
	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")
}

func readFilesV2() <-chan FileInfo {
	chanOut := make(chan FileInfo)

	go func() {
		err := filepath.WalkDir(tempPath, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}

			buf, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			chanOut <- FileInfo{
				FilePath: path,
				Content:  buf,
			}
			return nil
		})
		if err != nil {
			log.Println("ERROR:", err.Error())
		}
		close(chanOut)
	}()

	return chanOut
}

func getSum(chanIn <-chan FileInfo) <-chan FileInfo {
	chanOut := make(chan FileInfo)
	go func() {
		for fileInfo := range chanIn {
			fileInfo.Sum = fmt.Sprintf("%x", md5.Sum(fileInfo.Content))
			chanOut <- fileInfo
		}
		close(chanOut)
	}()
	return chanOut
}

func rename(chanIn <-chan FileInfo) <-chan FileInfo {
	chanOut := make(chan FileInfo)
	go func() {
		for fileInfo := range chanIn {
			newPath := filepath.Join(tempPath, fmt.Sprintf("file-%s.txt", fileInfo.Sum))
			err := os.Rename(fileInfo.FilePath, newPath)
			fileInfo.IsRenamed = err == nil
			chanOut <- fileInfo
		}
		close(chanOut)
	}()
	return chanOut
}

func mergeChanFileInfo(chanInMany ...<-chan FileInfo) <-chan FileInfo {
	wg := new(sync.WaitGroup)
	chanOut := make(chan FileInfo)

	wg.Add(len(chanInMany))
	for _, eachChan := range chanInMany {
		go func(eachChan <-chan FileInfo) {
			for eachChanData := range eachChan {
				chanOut <- eachChanData
			}
			wg.Done()
		}(eachChan)
	}

	go func() {
		wg.Wait()
		close(chanOut)
	}()

	return chanOut
}
