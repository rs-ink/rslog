package rslog

import (
	"bytes"
	"context"
	"runtime"
	"sync"
	"time"
)

type RsLog struct {
	Conf RsLoggerConfig
	mu   sync.Mutex
	buf  bytes.Buffer
	ctx  context.Context
}

var DefaultRsLog *RsLog

func init() {
	DefaultRsLog = NewRsLog(true)
}

func (r *RsLog) C(ctx context.Context) *RsLog {
	return &RsLog{
		Conf: r.Conf,
		mu:   sync.Mutex{},
		buf:  bytes.Buffer{},
		ctx:  ctx,
	}
}

type RsLoggerErrorGetter interface {
	GetError() error
}

func AssertError(err error, throwError ...error) {
	DefaultRsLog.AssertError(err, throwError...)
}

func (r RsLog) AssertError(err error, throwError ...error) {
	if err != nil {
		r.Out(1, LevelERROR, err)
		if len(throwError) > 0 && throwError[0] != nil {
			panic(throwError[0])
		} else {
			panic(err)
		}
	}
}

func C(ctx context.Context) *RsLog {
	return DefaultRsLog.C(ctx)
}

func NewRsLog(direct ...bool) *RsLog {
	conf := NewRConf(direct...)
	return &RsLog{
		Conf: conf,
	}
}

func SetProjectName(name string) {
	DefaultRsLog.SetProjectName(name)
}

func SetRLevel(level RLevel) {
	DefaultRsLog.Conf.SetRLevel(level, 1)
}

func SetRootRLevel(level RLevel) {
	DefaultRsLog.SetRootRLevel(level)
}

func Debug(v ...interface{}) {
	DefaultRsLog.Out(1, LevelDEBUG, v...)
}

func Info(v ...interface{}) {
	DefaultRsLog.Out(1, LevelINFO, v...)
}

func Warn(v ...interface{}) {
	DefaultRsLog.Out(1, LevelWARN, v...)
}

func Error(v ...interface{}) {
	DefaultRsLog.Out(1, LevelERROR, v...)
}

func DebugF(f string, v ...interface{}) {
	DefaultRsLog.OutF(1, LevelDEBUG, f, v...)
}

func InfoF(f string, v ...interface{}) {
	DefaultRsLog.OutF(1, LevelINFO, f, v...)
}

func WarnF(f string, v ...interface{}) {
	DefaultRsLog.OutF(1, LevelWARN, f, v...)
}

func ErrorF(f string, v ...interface{}) {
	DefaultRsLog.OutF(1, LevelERROR, f, v...)
}

func Out(callDepth int, level RLevel, v ...interface{}) {
	DefaultRsLog.OutPc(GetPcInfo(callDepth+1, DefaultRsLog.Conf.GetProjectName(), DefaultRsLog.Conf.IsDirect()), level, v...)
}

func OutF(callDepth int, level RLevel, f string, v ...interface{}) {
	DefaultRsLog.OutPcF(GetPcInfo(callDepth+1, DefaultRsLog.Conf.GetProjectName(), DefaultRsLog.Conf.IsDirect()), level, f, v...)
}
func OutPcF(pcInfo PcInfo, level RLevel, f string, v ...interface{}) {
	DefaultRsLog.OutPcF(pcInfo, level, f, v...)
}

func OutPc(pcInfo PcInfo, level RLevel, v ...interface{}) {
	DefaultRsLog.OutPc(pcInfo, level, v...)
}

func (r *RsLog) SetProjectName(name string) {
	r.Conf.SetProjectName(name)
}

func (r *RsLog) SetRLevel(level RLevel) {
	r.Conf.SetRLevel(level, 1)
}

func (r *RsLog) SetRootRLevel(level RLevel) {
	r.Conf.SetRootRLevel(level)
}

func (r *RsLog) Debug(v ...interface{}) {
	r.Out(1, LevelDEBUG, v...)
}

func (r *RsLog) Info(v ...interface{}) {
	r.Out(1, LevelINFO, v...)
}

func (r *RsLog) Warn(v ...interface{}) {
	r.Out(1, LevelWARN, v...)
}

func (r *RsLog) Error(v ...interface{}) {
	r.Out(1, LevelERROR, v...)
}

func (r *RsLog) DebugF(f string, v ...interface{}) {
	r.OutF(1, LevelDEBUG, f, v...)
}

func (r *RsLog) InfoF(f string, v ...interface{}) {
	r.OutF(1, LevelINFO, f, v...)
}

func (r *RsLog) WarnF(f string, v ...interface{}) {
	r.OutF(1, LevelWARN, f, v...)
}

func (r *RsLog) ErrorF(f string, v ...interface{}) {
	r.OutF(1, LevelERROR, f, v...)
}

func (r *RsLog) Out(callDepth int, level RLevel, v ...interface{}) {
	r.OutPc(GetPcInfo(callDepth+1, r.Conf.GetProjectName(), r.Conf.IsDirect()), level, v...)
}

func (r *RsLog) OutF(callDepth int, level RLevel, f string, v ...interface{}) {
	r.OutPcF(GetPcInfo(callDepth+1, r.Conf.GetProjectName(), r.Conf.IsDirect()), level, f, v...)
}

var lineSeparator string

func init() {
	lineSeparator = "\r\n"
	switch runtime.GOOS {
	case "windows":
		lineSeparator = "\r\n"
	case "linux":
		lineSeparator = "\n"
	}
}

func (r *RsLog) OutPc(pcInfo PcInfo, level RLevel, v ...interface{}) {
	targetLevel := r.Conf.GetRLevelPc(pcInfo)
	if level >= targetLevel {
		r.mu.Lock()
		defer r.mu.Unlock()
		defer r.buf.Reset()
		formatHeader(&r.buf, time.Now(), pcInfo, r.Conf.GetProjectName(), level.String())
		r.buf.WriteString(defaultRLogFormatWithContext(r.ctx, v...))
		r.buf.WriteString(lineSeparator)
		Writer(r.Conf.GetWriter(level), r.buf)
	}
}

func (r *RsLog) OutPcF(pcInfo PcInfo, level RLevel, f string, v ...interface{}) {
	targetLevel := r.Conf.GetRLevelPc(pcInfo)
	if level >= targetLevel {
		r.mu.Lock()
		defer r.mu.Unlock()
		defer r.buf.Reset()
		formatHeader(&r.buf, time.Now(), pcInfo, r.Conf.GetProjectName(), level.String())
		r.buf.WriteString(defaultRLogFormatFWithContext(r.ctx, f, v...))
		r.buf.WriteString(lineSeparator)
		Writer(r.Conf.GetWriter(level), r.buf)
	}
}

func (r *RsLog) SetRsLoggerConf(conf RsLoggerConfig) {
	r.Conf = conf
}
