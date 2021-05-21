package tbcrawler

import "testing"

func TestTextBookCrawler(t *testing.T) {
	// analyzeDirectory([]string {"C:\\Guru\\LearningMaterials"})
	analyzeDirectory("C:\\Guru\\LearningMaterials")
}