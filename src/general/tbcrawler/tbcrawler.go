package tbcrawler

import (
	"os"
	"log"
	"time"
)

/* 
type Pipeline struct {
	dirCh chan string
	filech chan string
	fileList []os.DirEntry
	rootDir string
}

 */
 func analyzeDirectory(dirs string) {
	if dirs == "" {
		return
	}

	stop := make(chan bool)
	var cnt int
	
	processDirs(dirs, stop, &cnt)

	for x := range stop {
		if x == false {
			log.Fatalf("error..... Quitting")
		}
	}

	log.Printf("Counted for (%q): %v \n", dirs, cnt)
}

func processDirs(dirs string, stop chan bool, cnt *int) int {
	var c int
	files, err := os.ReadDir(dirs)
	if err != nil {
		stop <- false
		close(stop)
		return c
	}

	t := time.NewTimer(500 * time.Millisecond)

	oLoop: for {
		select {
			default:
				for _, file := range files {
					if file.IsDir() {
						analyzeDirectory(dirs + "/" + file.Name())
					} else {
						log.Println(file.Name())
						*cnt++
						c++
					}
				}
			case <- t.C:
				close(stop)
				break oLoop
		}
	}
	return c
}
