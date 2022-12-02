package main

import (
	"testing"

	"github.com/YWJSonic/ycore/analysis/goanalysis"
	"github.com/YWJSonic/ycore/driver/load/project/goloader"
)

func TestGolandAnalysis(t *testing.T) {
	projectRootNode := goloader.LoadRoot("./")

	goanalysis.GoAnalysisSpaceFirst(projectRootNode)
}
