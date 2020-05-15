package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/novalagung/gubrak"
)

const totalFile = 3000
const contentLength = 5000
const totalWorker = 10
const timeoutDuration = 3 * time.Second

var tempPath = filepath.Join(os.Getenv("TEMP"), "chapter-A.59-pipeline-temp")

func main() {
	log.Println("start")
	start := time.Now()

	ctx, cancel := context.WithCancel(context.Background())
	time.AfterFunc(timeoutDuration, cancel)
	generate(ctx)

	// ctx, _ := context.WithTimeout(context.Background(), timeoutDuration)
	// generate(ctx)

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")
}

func generate(ctx context.Context) {
	os.RemoveAll(tempPath)
	os.MkdirAll(tempPath, os.ModePerm)

	jobs := make(chan int)

	wg := new(sync.WaitGroup)
	wg.Add(totalWorker)

	go dispatchWorkers(wg, jobs)
	go distributeJobs(jobs)

	done := make(chan bool)
	go func() {
		wg.Wait()
		done <- true
	}()

	select {
	case <-ctx.Done():
		log.Printf("generation process stopped. %s", ctx.Err())
	case <-done:
		log.Printf("%d of total files created", totalFile)
	}
}

func startWorker(wg *sync.WaitGroup, jobs <-chan int, workerNumber int) {
	log.Printf("worker-%d started", workerNumber)

	i := 0

	for jobNumber := range jobs {
		filename := filepath.Join(tempPath, fmt.Sprintf("file-%d.txt", jobNumber))
		content := gubrak.RandomString(contentLength)
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
