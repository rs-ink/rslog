package example

import (
	"context"
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

func TestWithContext(t *testing.T) {
	ctx := context.WithValue(context.Background(), "ip", "127.0.0.1")
	ctx = context.WithValue(ctx, "requestId", "123123")
	rslog.RegisterRLogFormatContextKey([]string{"requestId", "ips"})
	l := rslog.C(ctx)
	l.Info("asdfasdfasdfasdf")
}

func BenchmarkRsLog(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rslog.Debug("asdfasdfasdf")
	}
}
