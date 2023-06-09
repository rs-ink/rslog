# rslog

## 2023-06-09 
* 新增扩展内容支持
* 优化调用逻辑

## starter 项目中使用
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
