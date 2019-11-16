package example

import (
	"rslog"
	"testing"
)

var log *rslog.RsLog

func init() {
	log = rslog.DefaultRsLog
	log.SetProjectName("rslog")
	log.SetRootRLevel(rslog.LevelDEBUG)
	log.SetRLevel(rslog.LevelINFO)
}

func TestRsLog(t *testing.T) {
	log.Warn("asdfasdfasdf")
}

func BenchmarkRsLog(b *testing.B) {

	for i := 0; i < b.N; i++ {
		log.Debug("asdfasdfasdf")
	}
}
