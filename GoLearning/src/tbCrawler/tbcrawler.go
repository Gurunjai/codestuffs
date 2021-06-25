package tbcrawler

import "fmt"

func analyzeDirectory(dirs []string) {
	if 0 > len(dirs) {
		return
	}

	fmt.Printf("Got Dirs: %+v\n", dirs)
}