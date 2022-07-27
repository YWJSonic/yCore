package goanalysis

import (
	"os"
	"strings"
	"ycore/driver/load/project/goloader"
	"ycore/util"
)

// 專案節點樹分析 廣度優先
func GoAnalysisSpaceFirst(node *goloader.GoFileNode) {
	var childs []*goloader.GoFileNode

	// 取得節點下的全部子節點
	getSubChiles := func(node *goloader.GoFileNode) []*goloader.GoFileNode {
		var childs []*goloader.GoFileNode
		for _, child := range node.Childes {
			if child.Childes != nil {
				childs = append(childs, child.Childes...)
			}
		}
		return childs
	}

	childs = node.Childes
	var currentChild *goloader.GoFileNode
	// 要輪尋的節點尚未
	for len(childs) > 0 {

		currentChild, childs = childs[0], childs[1:]

		if currentChild.Childes != nil { // 節點下還有子節點
			childs = append(childs, getSubChiles(currentChild)...)

		} else if currentChild.FileType == "go" { // 檔案節點
			// 讀檔
			codeLine := util.ReadFileToLineStr(currentChild.Path())
			// 檔案分析
			GoCodeAnalysis(codeLine)
		}
	}

}

// 專案節點樹分析  深度優先
func GoAnalysisDepthFirst(node *goloader.GoFileNode) {
	// 節點下還有子節點
	for _, child := range node.Childes {
		GoAnalysisDepthFirst(child)
	}

	if node.FileType == "go" { // 檔案節點
		// 讀檔
		codeLine := util.ReadFileToLineStr(node.Path())
		// 檔案分析
		GoCodeAnalysis(codeLine)
	}
}

// 語言分析
func GoCodeAnalysis(codeLine []string) {
	lineCount := -1
	funcBlock := false
	importBlock := false

	allPackageMap := map[string]*PackageInfo{}
	var currentPackage *PackageInfo
	// var currentFunc *FuncInfo
	var funcBlockCount uint
	for {
		lineCount++
		if len(codeLine) == lineCount {
			break
		}
		currentLine := codeLine[lineCount]
		if importBlock { // 分析 import 檔
			if currentLine == ")" {
				importBlock = false
				continue
			}

			// 清除格式美化
			cleanReplacer := strings.NewReplacer("\"", "", "\t", "")
			currentLine = cleanReplacer.Replace(currentLine)

			// import 參數
			filePathSplit := strings.Split(currentLine, " ")
			newName := ""

			// package 有重新命名
			if len(filePathSplit) > 1 {
				newName = filePathSplit[0]
			}

			// 取得資料夾名稱
			filePathSplit = strings.Split(currentLine, string(os.PathSeparator))
			pacakageName := filePathSplit[len(filePathSplit)-1]

			// package 關聯建立
			if packageInfo, ok := allPackageMap[pacakageName]; ok {
				packageLink := NewPackageLink(newName, packageInfo)
				currentPackage.ImportLink[packageLink.Name()] = packageLink
				allPackageMap[packageInfo.Name] = packageInfo
			} else {
				packageInfo = NewPackageInfo(pacakageName)
				packageLink := NewPackageLink(newName, packageInfo)
				currentPackage.ImportLink[packageLink.Name()] = packageLink
				allPackageMap[packageInfo.Name] = packageInfo

			}

		} else if funcBlock { // 在 func 區塊內

			funcBlock = false
			splitStr := strings.Split(currentLine, "//")
			currentLine = splitStr[0]
			if currentLine[len(currentLine)-1] == '{' {
				funcBlockCount++
			} else if currentLine[len(currentLine)-1] == '}' {
				funcBlockCount--
			}

		} else { // 不再任何區塊內

			splitStr := strings.Split(currentLine, " ")

			if splitStr[0] == "package" { // 分析 package 命名
				packageInfo := NewPackageInfo(splitStr[1])
				allPackageMap[packageInfo.Name] = packageInfo
				currentPackage = packageInfo
			} else if splitStr[0] == "import" { // 進入 import 區塊

				if splitStr[1] == "(" {
					importBlock = true
				} else { // 單行 import 直接分析

					// 建立關聯
					filePathSplit := strings.Split(splitStr[1], string(os.PathSeparator)) // os不一樣方向不一樣
					pacakageName := filePathSplit[len(filePathSplit)-1]
					pacakgeInfo := NewPackageInfo(pacakageName)
					packageLink := NewPackageLink("", pacakgeInfo)
					currentPackage.ImportLink[packageLink.Name()] = packageLink
				}

			} else if splitStr[0] == "func" { // 進入方法區塊
				funcBlock = true
				funcBlockCount++
				// SplitFunc(currentLine)

				//建立關聯
				// currentFunc = NewFuncInfo(splitStr[1])

			}
		}
	}
}

func SplitFunc(codeLine string) []string {
	var splitStr []string

	return splitStr
}
