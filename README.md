# rslog

* 基于func路径log配置
```` 
func init() {
	//过滤显示项目名称
	rslog.SetProjectName("rslog")
	//设置根Level 
	rslog.SetRootRLevel(rslog.LevelDEBUG)
	设置pkg根Level 
	rslog.SetRLevel(rslog.LevelINFO)
}

func TestRsLog(t *testing.T) {
	rslog.Warn("asdfasdfasdf")
}
````
