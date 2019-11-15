package rslog

import (
	"bytes"
	"fmt"
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
	DefaultRsLog = &RsLog{
		Conf: NewRConf(),
	}
}

func SetProjectName(name string) {
	DefaultRsLog.Conf.SetProjectName(name)
}

func SetRLevel(level RLevel) {
	DefaultRsLog.Conf.SetRLevel(level, 1)
}

func SetRootRLevel(level RLevel) {
	DefaultRsLog.Conf.SetRootRLevel(level)
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

func (rl RsLog) Out(callDepth int, level RLevel, v ...interface{}) {
	rl.OutPc(GetPcInfo(callDepth+1, rl.Conf.GetProjectName()), level, v...)
}

func (rl RsLog) OutF(callDepth int, level RLevel, f string, v ...interface{}) {
	rl.OutPcF(GetPcInfo(callDepth+1, rl.Conf.GetProjectName()), level, f, v...)
}

func (rl RsLog) OutPc(pcInfo PcInfo, level RLevel, v ...interface{}) {
	targetLevel := rl.Conf.GetRLevelPc(pcInfo)
	if level >= targetLevel {
		rl.mu.Lock()
		defer rl.mu.Unlock()
		defer rl.buf.Reset()
		formatHeader(&rl.buf, time.Now(), pcInfo, rl.Conf.GetProjectName())
		rl.buf.WriteString("[")
		rl.buf.WriteString(level.String())
		rl.buf.WriteString("] ")
		rl.buf.WriteString(fmt.Sprint(v...))
		rl.buf.WriteString("\r\n")
		Writer(rl.Conf.GetWriter(level), rl.buf)
	}
}

func (rl RsLog) OutPcF(pcInfo PcInfo, level RLevel, f string, v ...interface{}) {
	targetLevel := rl.Conf.GetRLevelPc(pcInfo)
	if level >= targetLevel {
		fmt.Println(targetLevel)
		fmt.Println(level)

		rl.mu.Lock()
		defer rl.mu.Unlock()
		defer rl.buf.Reset()
		formatHeader(&rl.buf, time.Now(), pcInfo, rl.Conf.GetProjectName())
		rl.buf.WriteString("[")
		rl.buf.WriteString(level.String())
		rl.buf.WriteString("] ")
		rl.buf.WriteString(fmt.Sprintf(f, v...))
		rl.buf.WriteString("\r\n")
		Writer(rl.Conf.GetWriter(level), rl.buf)
	}
}

func (rl *RsLog) SetRsLoggerConf(conf RsLoggerConfig) {
	rl.Conf = conf
}
