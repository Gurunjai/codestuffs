package tbcrawler

import "os"
import "log"

type Pipeline struct {
	dirCh chan string
	filech chan string
	fileList []os.DirEntry
	rootDir string
}

func NewPipeLine()
func analyzeDirectory(dirs string) {
	if dirs == "" {
		return
	}

	dirCh := make(chan string)
	stop := make(chan bool)

	go processDirs(dirCh, stop)
	dirCh <- dirs
	for x := range stop {
		if x == false {
			log.Fatalf("error..... Quitting")
		}
	}
}

func processDirs(dirCh chan string, stop chan bool) {
	for out := range dirCh {
		files, err := os.ReadDir(out)
		if err != nil {
			log.Printf("Failed to process the directory")
			stop <- false
		}
		
		for _, file := range files {
			if file.IsDir() {
				dirCh <- file.Name()
			}

			log.Println(file.Name())
		}
	}
}