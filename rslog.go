package rslog

import (
	"bytes"
	"fmt"
	"os"
	"sync"
	"time"
)

type RsLog struct {
	Conf RsLoggerConfig
	mu   sync.Mutex
	buf  bytes.Buffer
}

var DefaultRsLog *RsLog

func init() {
	DefaultRsLog = NewRsLog(true)
}

func NewRsLog(direct ...bool) *RsLog {
	return &RsLog{
		Conf: NewRConf(direct...),
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

func (rl *RsLog) SetProjectName(name string) {
	rl.Conf.SetProjectName(name)
}

func (rl *RsLog) SetRLevel(level RLevel) {
	rl.Conf.SetRLevel(level, 1)
}

func (rl *RsLog) SetRootRLevel(level RLevel) {
	rl.Conf.SetRootRLevel(level)
}

func (rl *RsLog) Debug(v ...interface{}) {
	rl.Out(1, LevelDEBUG, v...)
}

func (rl *RsLog) Info(v ...interface{}) {
	rl.Out(1, LevelINFO, v...)
}

func (rl *RsLog) Warn(v ...interface{}) {
	rl.Out(1, LevelWARN, v...)
}

func (rl *RsLog) Error(v ...interface{}) {
	rl.Out(1, LevelERROR, v...)
}

func (rl *RsLog) DebugF(f string, v ...interface{}) {
	rl.OutF(1, LevelDEBUG, f, v...)
}

func (rl *RsLog) InfoF(f string, v ...interface{}) {
	rl.OutF(1, LevelINFO, f, v...)
}

func (rl *RsLog) WarnF(f string, v ...interface{}) {
	rl.OutF(1, LevelWARN, f, v...)
}

func (rl *RsLog) ErrorF(f string, v ...interface{}) {
	rl.OutF(1, LevelERROR, f, v...)
}

func (rl *RsLog) Out(callDepth int, level RLevel, v ...interface{}) {
	rl.OutPc(GetPcInfo(callDepth+1, rl.Conf.GetProjectName(), rl.Conf.IsDirect()), level, v...)
}

func (rl *RsLog) OutF(callDepth int, level RLevel, f string, v ...interface{}) {
	rl.OutPcF(GetPcInfo(callDepth+1, rl.Conf.GetProjectName(), rl.Conf.IsDirect()), level, f, v...)
}

var lineSeparator string

func init() {
	lineSeparator = os.Getenv("line.separator")
}

func (rl *RsLog) OutPc(pcInfo PcInfo, level RLevel, v ...interface{}) {
	targetLevel := rl.Conf.GetRLevelPc(pcInfo)
	if level >= targetLevel {
		rl.mu.Lock()
		defer rl.mu.Unlock()
		defer rl.buf.Reset()
		formatHeader(&rl.buf, time.Now(), pcInfo, rl.Conf.GetProjectName(), level.String())
		rl.buf.WriteString(fmt.Sprint(v...))
		rl.buf.WriteString(lineSeparator)
		Writer(rl.Conf.GetWriter(level), rl.buf)
	}
}

func (rl *RsLog) OutPcF(pcInfo PcInfo, level RLevel, f string, v ...interface{}) {
	targetLevel := rl.Conf.GetRLevelPc(pcInfo)
	if level >= targetLevel {
		rl.mu.Lock()
		defer rl.mu.Unlock()
		defer rl.buf.Reset()
		formatHeader(&rl.buf, time.Now(), pcInfo, rl.Conf.GetProjectName(), level.String())
		rl.buf.WriteString(fmt.Sprintf(f, v...))
		rl.buf.WriteString(lineSeparator)
		Writer(rl.Conf.GetWriter(level), rl.buf)
	}
}

func (rl *RsLog) SetRsLoggerConf(conf RsLoggerConfig) {
	rl.Conf = conf
}
