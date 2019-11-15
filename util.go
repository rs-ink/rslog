package rslog

import (
	"bytes"
	"io"
	"runtime"
	"strconv"
	"strings"
)

const funcSep = "."

func GetFuncName(info PcInfo, projectName string) string {
	fName := runtime.FuncForPC(info.Pc).Name()
	if projectName == "" {
		return fName
	}
	if strings.Contains(fName, projectName+"/") {
		return strings.Replace(strings.Split(fName, projectName+"/")[1], "/", funcSep, -1)
	} else {
		return strings.Replace(fName, "/", funcSep, -1)
	}
}

func GetRealFuncName(pcInfo PcInfo, projectName string) []string {
	funcName := GetFuncName(pcInfo, projectName)
	fs := strings.Split(funcName, funcSep)
	_, err := strconv.ParseInt(fs[len(fs)-1], 10, 64)
	if err == nil {
		return fs[:len(fs)-1]
	}
	return fs
}

func Writer(writer io.Writer, buf bytes.Buffer) {
	_, _ = writer.Write(buf.Bytes())
}
