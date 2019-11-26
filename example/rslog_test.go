package example

import (
	"github.com/rs-ink/rslog"
	"testing"
)

func init() {
	rslog.SetProjectName("rslog")
	rslog.SetRootRLevel(rslog.LevelDEBUG)
	rslog.SetRLevel(rslog.LevelINFO)
}

func TestRsLog(t *testing.T) {
	rslog.Warn("asdfasdfasdf")
}

func BenchmarkRsLog(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rslog.Debug("asdfasdfasdf")
	}
}
