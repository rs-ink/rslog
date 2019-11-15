# rslog

* 基于func路径log配置
```` 
func init() {
	//过滤显示项目名称
	rslog.DefaultRsLog.Conf.SetProjectName("rslog")
	//设置根Level 
	rslog.DefaultRsLog.Conf.SetRootRLevel(rslog.LevelDEBUG)
	设置pkg根Level 
	rslog.DefaultRsLog.Conf.SetRLevel(rslog.LevelINFO)
}

func TestRsLog(t *testing.T) {
	rslog.Warn("asdfasdfasdf")
}
````
