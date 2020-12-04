package example

import (
	"github.com/rs-ink/rslog"
	"io/ioutil"
	"os"
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

var buildArgs []string
var _ = initBuild()

func initBuild() bool {
	buildArgs = os.Args
	return true
}

func TestRsLog(t *testing.T) {
	//ps := strings.Split(os.Getenv("GOMOD"), string(os.PathSeparator))
	envs := os.Environ()
	for _, env := range envs {
		rslog.Warn(env)
	}
	rslog.Warn("===========================")
	for _, arg := range os.Args {
		rslog.Warn(arg)
	}
	rslog.Warn("===========================")
	for _, arg := range buildArgs {
		rslog.Warn(arg)
	}
}

func BenchmarkRsLog(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rslog.Debug("asdfasdfasdf")
	}
}
