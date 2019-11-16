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

func (rl RsLog) SetProjectName(name string) {
	rl.Conf.SetProjectName(name)
}

func (rl RsLog) SetRLevel(level RLevel) {
	rl.Conf.SetRLevel(level, 1)
}

func (rl RsLog) SetRootRLevel(level RLevel) {
	rl.Conf.SetRootRLevel(level)
}

func (rl RsLog) Debug(v ...interface{}) {
	rl.Out(1, LevelDEBUG, v...)
}

func (rl RsLog) Info(v ...interface{}) {
	rl.Out(1, LevelINFO, v...)
}

func (rl RsLog) Warn(v ...interface{}) {
	rl.Out(1, LevelWARN, v...)
}

func (rl RsLog) Error(v ...interface{}) {
	rl.Out(1, LevelERROR, v...)
}

func (rl RsLog) DebugF(f string, v ...interface{}) {
	rl.OutF(1, LevelDEBUG, f, v...)
}

func (rl RsLog) InfoF(f string, v ...interface{}) {
	rl.OutF(1, LevelINFO, f, v...)
}

func (rl RsLog) WarnF(f string, v ...interface{}) {
	rl.OutF(1, LevelWARN, f, v...)
}

func (rl RsLog) ErrorF(f string, v ...interface{}) {
	rl.OutF(1, LevelERROR, f, v...)
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
