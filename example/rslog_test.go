package example

import (
	"github.com/rs-ink/rslog"
	"io/ioutil"
	"os"
	"syscall"
	"testing"
	"time"
)

func init() {
	rslog.SetProjectName("rslog")
	rslog.SetRootRLevel(rslog.LevelDEBUG)
	rslog.SetRLevel(rslog.LevelINFO)
}

type writer struct {
}

func (w writer) Write(p []byte) (n int, err error) {
	return len(p), ioutil.WriteFile(time.Now().Format("./logs/20060102150405.log"), p, os.ModeAppend)
}

func TestRsLog(t *testing.T) {
	envs := syscall.Environ()
	for _, env := range envs {
		rslog.Warn(env)
	}

}

func BenchmarkRsLog(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rslog.Debug("asdfasdfasdf")
	}
}
