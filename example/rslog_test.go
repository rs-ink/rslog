package example

import (
	"rslog"
	"testing"
)

func init() {
	rslog.DefaultRsLog.Conf.SetProjectName("rslog")
	rslog.DefaultRsLog.Conf.SetRootRLevel(rslog.LevelDEBUG)
	rslog.DefaultRsLog.Conf.SetRLevel(rslog.LevelINFO, 0)
}

func TestRsLog(t *testing.T) {
	rslog.Warn("asdfasdfasdf")
}

func BenchmarkRsLog(b *testing.B) {

	for i := 0; i < b.N; i++ {
		rslog.Debug("asdfasdfasdf")
	}
}
