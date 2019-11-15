package rslog

import (
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

type RConf struct {
	DefaultWriter io.Writer
	MapWriter     map[RLevel]io.Writer
	Debug         bool
	RootLevel     RLevel
	ProjectName   string
	mu            sync.Mutex
	MapLevel      *sync.Map
}

func (rc RConf) GetRLevelPc(info PcInfo) RLevel {
	funcNames := GetFuncName(info, rc.ProjectName)
	fs := strings.Split(funcNames, funcSep)
	var i int
	for i = len(fs); i >= 1; i-- {
		if l, ok := rc.MapLevel.Load(strings.ToLower(strings.Join(fs[:i], funcSep))); ok {
			return l.(RLevel)
		}
	}

	if rc.RootLevel == 0 {
		return LevelDEBUG
	} else {
		return rc.RootLevel
	}
}

//代码段设置Level
func (rc RConf) SetRLevel(level RLevel, callDepth int) {
	pcInfo := GetPcInfo(callDepth+1, rc.ProjectName)
	funcNames := GetRealFuncName(pcInfo, rc.ProjectName)
	if funcNames[len(funcNames)-1] == "init" {
		funcNames = funcNames[:len(funcNames)-1]
	}
	rc.MapLevel.LoadOrStore(strings.ToLower(strings.Join(funcNames, funcSep)), level)
}

func (rc RConf) GetRLevel(callDepth int) RLevel {
	return rc.GetRLevelPc(GetPcInfo(callDepth+1, rc.ProjectName))
}

func NewRConf() *RConf {
	return &RConf{
		MapWriter: make(map[RLevel]io.Writer),
		MapLevel:  &sync.Map{},
	}
}

func (rc *RConf) SetWriter(writer io.Writer) {
	rc.DefaultWriter = writer
}

func (rc *RConf) SetWriterForLevel(level RLevel, writer io.Writer) {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	if rc.MapWriter == nil {
		rc.MapWriter = make(map[RLevel]io.Writer)
	}
	rc.MapWriter[level] = writer
}

func (rc *RConf) SetProjectName(name string) {
	rc.ProjectName = name
}

func (rc RConf) GetProjectName() string {
	return rc.ProjectName
}

func (rc RConf) GetWriter(level RLevel) io.Writer {
	if w, ok := rc.MapWriter[level]; ok {
		return w
	} else if rc.DefaultWriter != nil {
		return rc.DefaultWriter
	} else {
		if level >= LevelWARN {
			return os.Stderr
		} else {
			return os.Stdout
		}
	}
}

func (rc *RConf) SetRootRLevel(level RLevel) {
	rc.RootLevel = level
}

func (rc RConf) IsDebug() bool {
	return rc.Debug
}

func (rc *RConf) ScanConfDuration(duration time.Duration) {
	//TODO 扫描配置文件更新配置 异步
}
