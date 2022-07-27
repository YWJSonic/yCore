package myyaml

import "ycore/load/file/yamlloader"

func Load(path string, Params interface{}) error {
	// 讀取遊戲設定
	return yamlloader.LoadYaml(path, Params)
}
