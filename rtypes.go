package rslog

import (
	"io"
	"os"
	"runtime"
	"strings"
	"time"
)

type RLevel int

const (
	_ RLevel = iota
	LevelDEBUG
	LevelINFO
	LevelWARN
	LevelERROR
	LevelOFF
)

func (level RLevel) String() string {
	switch level {
	case
		LevelDEBUG:
		return "DEBUG"
	case LevelINFO:
		return "INFO"
	case LevelWARN:
		return "WARN"
	case LevelERROR:
		return "ERROR"
	case LevelOFF:
		return "OFF"
	default:
		return ""
	}
}

//指定显示路径页面info
type PcInfo struct {
	Pc   uintptr
	File string
	Line int
	Ok   bool
}

var pkgCache = os.Getenv("GOMODCACHE")
var goROOT = os.Getenv("GOROOT")
var goPath = os.Getenv("GOPATH")

func GetPcInfo(callDepth int, projectName string, direct ...bool) (pcInfo PcInfo) {
	for i := 1; i < 10; i++ {
		pcInfo.Pc, pcInfo.File, pcInfo.Line, pcInfo.Ok = runtime.Caller(callDepth + i)
		if len(direct) > 0 && direct[0] {
			return
		}
		if projectName == "" {
			return
		}
		if strings.Contains(pcInfo.File, projectName) && !strings.HasPrefix(pcInfo.File, pkgCache) && !strings.HasPrefix(pcInfo.File, goROOT) && !strings.HasPrefix(pcInfo.File, goPath) {
			return
		}
	}
	return
}

type RsLoggerConfig interface {
	SetWriter(writer io.Writer)
	SetWriterForLevel(level RLevel, writer io.Writer)
	SetProjectName(name string)
	GetProjectName() string
	GetWriter(level RLevel) io.Writer
	SetRootRLevel(level RLevel)
	SetRLevel(level RLevel, callDepth int)
	GetRLevelPc(info PcInfo) RLevel
	GetRLevel(callDepth int) RLevel
	IsDebug() bool
	ScanConfDuration(duration time.Duration)
	IsDirect() bool
	SetDirect(direct bool)
}

type RsLogger interface {
	Out(callDepth int, level RLevel, v ...interface{})
	OutF(callDepth int, level RLevel, f string, v ...interface{})
	OutPc(pcInfo PcInfo, level RLevel, v ...interface{})
	OutPcF(pcInfo PcInfo, level RLevel, f string, v ...interface{})
	SetRsLoggerConf(conf RsLoggerConfig)
}

type RsLoggerIProject interface {
	SetProjectName(name string)
}

type RsLoggerILevel interface {
	SetRLevel(level RLevel)
	SetRootRLevel(level RLevel)
}

type RsLoggerI interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
	DebugF(f string, v ...interface{})
	InfoF(f string, v ...interface{})
	WarnF(f string, v ...interface{})
	ErrorF(f string, v ...interface{})
}
