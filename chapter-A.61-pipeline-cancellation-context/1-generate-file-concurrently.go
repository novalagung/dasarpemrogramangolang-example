package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const totalFile = 3000
const contentLength = 5000
const totalWorker = 10

var tempPath = filepath.Join(os.Getenv("TEMP"), "chapter-A.61-pipeline-cancellation-context")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	log.Println("start")
	start := time.Now()

	generateFiles()

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")
}

func randomString(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func generateFiles() {
	os.RemoveAll(tempPath)
	os.MkdirAll(tempPath, os.ModePerm)

	jobs := make(chan int)

	wg := new(sync.WaitGroup)
	wg.Add(totalWorker)

	go dispatchWorkers(wg, jobs)
	go distributeJobs(jobs)

	wg.Wait()

	log.Printf("%d of total files created", totalFile)
}

func startWorker(wg *sync.WaitGroup, jobs <-chan int, workerNumber int) {
	log.Printf("worker-%d started", workerNumber)

	i := 0

	for jobNumber := range jobs {
		filename := filepath.Join(tempPath, fmt.Sprintf("file-%d.txt", jobNumber))
		content := randomString(contentLength)
		err := ioutil.WriteFile(filename, []byte(content), os.ModePerm)
		if err != nil {
			log.Println("Error writing file", filename)
		}

		if i > 0 && i%100 == 0 {
			log.Printf(" -> worker-%d created %d files", workerNumber, i)
		}

		i++
	}

	log.Printf(" -> worker-%d done creating %d files", workerNumber, i)
	wg.Done()
}

func dispatchWorkers(wg *sync.WaitGroup, jobs chan int) {
	for i := 0; i < totalWorker; i++ {
		go startWorker(wg, jobs, i)
	}
}

func distributeJobs(jobs chan int) {
	for i := 0; i < totalFile; i++ {
		jobs <- i
	}
	close(jobs)
}
