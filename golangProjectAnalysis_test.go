package main

import (
	"testing"
	"ycore/analysis/goanalysis"
	"ycore/load/project/goloader"
)

func TestGolandAnalysis(t *testing.T) {
	projectRootNode := goloader.LoadRoot("./")

	goanalysis.GoAnalysisSpaceFirst(projectRootNode)
}
