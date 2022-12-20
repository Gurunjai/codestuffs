package tbcrawler

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

/*
type Pipeline struct {
	dirCh chan string
	fileCh chan string
	fileList []os.DirEntry
	rootDir string
}

*/

type Crawl struct {
	pLog  *log.Logger
	dChan chan string
	fChan chan string
	stop  chan bool
	wg    *sync.WaitGroup
}

func NewCrawl(file *os.File) *Crawl {
	return &Crawl{
		pLog:  log.New(file, "", 0),
		dChan: make(chan string),
		fChan: make(chan string),
		stop:  make(chan bool),
		wg:    &sync.WaitGroup{},
	}
}

func dirCrawler(dirs []string) {
	if 0 >= len(dirs) {
		return
	}

	currentTime := time.Now()
	outFile := fmt.Sprintf("../out/materials_found-%d%d%d-%d:%d:%d.txt",
		currentTime.Year(), currentTime.Month(), currentTime.Day(),
		currentTime.Hour(), currentTime.Minute(), currentTime.Second())
	file, err := os.OpenFile(outFile, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	ca := NewCrawl(file)

	ca.analyzeDirectory()
	for _, dir := range dirs {
		ca.processInput(dir)
	}

	ca.wg.Wait()
	close(ca.dChan)
	close(ca.fChan)
	close(ca.stop)
}

func (ca *Crawl) processInput(dir string) {
	ca.processDirs(dir)
}

func (c *Crawl) analyzeDirectory() {
	go func() {
	oLoop:
		for {
			select {
			case dir := <-c.dChan:
				// c.pLog.Printf("Processing for Directory: %v\n", dir)
				go c.processDirs(dir)
			case file := <-c.fChan:
				if strings.HasSuffix(file, ".jpg") || strings.Contains(file, ".DS_Store") {
					continue oLoop
				}
				c.pLog.Println(file)
			case <-c.stop:
				c.pLog.Printf("Program Terminated...")
				break oLoop
			}
		}
	}()
}

func (ca *Crawl) processDirs(dirs string) {
	ca.wg.Add(1)
	go func(dirs string) {
		defer ca.wg.Done()
		if dirs == "" {
			return
		}

		files, err := os.ReadDir(dirs)
		if err != nil {
			ca.pLog.Printf("Error: %v opening directory: %q\n", err.Error(), dirs)
			return
		}

		for _, file := range files {
			if file.IsDir() {
				ca.dChan <- dirs + "/" + file.Name()
			} else {
				strTmp := fmt.Sprintf("%s:\n%+200s", dirs, file.Name())
				ca.fChan <- strTmp
			}
		}
	}(dirs)
}
