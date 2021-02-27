package main

import (
	"testing"
	"yangServer/analysis/goanalysis"
	"yangServer/load/project/goloader"
)

func TestGolandAnalysis(t *testing.T) {
	projectRootNode := goloader.LoadRoot("./")

	goanalysis.GoAnalysisSpaceFirst(projectRootNode)
}
